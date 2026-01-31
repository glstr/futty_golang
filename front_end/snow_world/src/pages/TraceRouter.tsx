import { useState } from 'react'
import './TraceRouter.css'
import { apiConfig } from '../config/api'

interface TraceResult {
  ttl: number
  network: string
  addr: string
  duration: string
  error?: string
  country?: string
  region?: string
  city?: string
  isp?: string
}

function TraceRouter() {
  const [ipAddress, setIpAddress] = useState('')
  const [results, setResults] = useState<TraceResult[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleTrace = async () => {
    if (!ipAddress.trim()) {
      setError('请输入 IP 地址或域名（例如：8.8.8.8 或 www.example.com）')
      return
    }

    setIsLoading(true)
    setError(null)
    setResults([])

    try {
      const response = await fetch(`${apiConfig.baseURL}/snow/router/trace`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ip: ipAddress.trim(),
        }),
      })

      const data = await response.json()

      if (!response.ok || data.error_code !== 0) {
        throw new Error(data.error_msg || '请求失败')
      }

      // 处理返回结果
      if (data.task_id) {
        // 如果是异步任务，轮询获取结果
        pollTaskResult(data.task_id)
      } else {
        setError('未收到任务ID')
      }
    } catch (err) {
      console.error('Trace router error:', err)
      setError(err instanceof Error ? err.message : '未知错误')
    } finally {
      setIsLoading(false)
    }
  }

  const pollTaskResult = async (taskId: number) => {
    const maxAttempts = 60 // 最多轮询60次
    let attempts = 0

    const poll = async () => {
      try {
        const response = await fetch(`${apiConfig.baseURL}/snow/router/get`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task_id: taskId,
          }),
        })

        const data = await response.json()

        if (data.error_code !== 0) {
          // error_code 不为 0，立即停止轮询
          setError(data.error_msg || '请求失败')
          setIsLoading(false)
          return
        }

        // 根据 state 决定是否继续轮询
        // state: 0 = 初始化, 1 = 完成, 2 = 失败
        const state = data.state !== undefined ? data.state : 0

        if (state === 0) {
          // 状态为0（初始化），继续轮询
          if (attempts < maxAttempts) {
            attempts++
            setTimeout(poll, 1000) // 继续轮询
          } else {
            setError('请求超时，请稍后重试')
            setIsLoading(false)
          }
          return
        }

        // state 为 1（完成）或 2（失败），停止轮询
        setIsLoading(false)

        if (state === 1) {
          // 任务完成，解析并显示结果
          if (data.result) {
            try {
              const parsed = JSON.parse(data.result)
              if (Array.isArray(parsed)) {
                formatAndSetResults(parsed)
              } else {
                // 如果不是数组，尝试其他格式
                formatAndSetResults([parsed])
              }
            } catch {
              // 如果不是JSON，直接显示文本
              setResults([{
                ttl: 0,
                network: '',
                addr: data.result,
                duration: '',
              }])
            }
          } else {
            setError('未收到有效结果')
          }
        } else if (state === 2) {
          // 任务失败
          setError('路由追踪失败，请稍后重试')
        }
      } catch (err) {
        console.error('Poll error:', err)
        if (attempts < maxAttempts) {
          attempts++
          setTimeout(poll, 1000)
        } else {
          setError('获取结果失败')
          setIsLoading(false)
        }
      }
    }

    poll()
  }

  const formatAndSetResults = (rawResults: any[]) => {
    const formatted: TraceResult[] = rawResults.map((result: any) => ({
      ttl: result.ttl || result.TTL || 0,
      network: result.network || result.Network || '',
      addr: result.addr || result.Addr || '',
      duration: result.duration || result.Duration || '0ms',
      error: result.error || (result.Error ? result.Error.toString() : undefined),
      country: result.country || result.Country || '',
      region: result.region || result.Region || '',
      city: result.city || result.City || '',
      isp: result.isp || result.ISP || '',
    }))
    setResults(formatted)
  }

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && !isLoading) {
      handleTrace()
    }
  }

  const handleClear = () => {
    setResults([])
    setError(null)
    setIpAddress('')
  }

  return (
    <div className="trace-router-container">
      <div className="trace-router-header">
        <h1>路由追踪工具</h1>
        <p className="trace-router-subtitle">输入IP地址或者域名，追踪数据包经过的路由节点</p>
      </div>

      <div className="trace-router-input-section">
        <div className="input-wrapper">
          <input
            type="text"
            className="trace-input"
            value={ipAddress}
            onChange={(e) => setIpAddress(e.target.value)}
            onKeyPress={handleKeyPress}
            placeholder="请输入IP地址或者域名（如：8.8.8.8）"
            disabled={isLoading}
          />
          <div className="button-group">
            <button
              className="trace-button"
              onClick={handleTrace}
              disabled={isLoading || !ipAddress.trim()}
            >
              {isLoading ? '追踪中...' : '开始追踪'}
            </button>
            <button
              className="clear-button"
              onClick={handleClear}
              disabled={isLoading}
            >
              清空
            </button>
          </div>
        </div>
        {error && (
          <div className="error-message">
            {error}
          </div>
        )}
      </div>

      <div className="trace-router-results">
        <div className="results-header">
          <h2>追踪结果</h2>
          {results.length > 0 && (
            <span className="results-count">共 {results.length} 跳</span>
          )}
        </div>
        <div className="results-content">
          {isLoading && results.length === 0 && (
            <div className="loading-message">
              <div className="loading-spinner"></div>
              <span>正在追踪路由，请稍候...</span>
            </div>
          )}
          {!isLoading && results.length === 0 && !error && (
            <div className="empty-message">
              请输入IP地址并点击"开始追踪"按钮
            </div>
          )}
          {results.length > 0 && (
            <div className="results-table">
              <div className="table-header">
                <div className="col-ttl">跳数</div>
                <div className="col-addr">IP地址</div>
                <div className="col-duration">延迟</div>
                <div className="col-status">状态</div>
              </div>
              {results.map((result, index) => (
                <div key={index} className="table-row">
                  <div className="col-ttl">{result.ttl}</div>
                  <div className="col-addr">
                    {result.addr || result.error || '*'}
                    {(result.country || result.region || result.city || result.isp) && (
                      <div className="addr-extra">
                        {[result.country, result.region, result.city].filter(Boolean).join(' / ')}
                        {result.isp ? ` · ${result.isp}` : ''}
                      </div>
                    )}
                  </div>
                  <div className="col-duration">
                    {result.duration || '-'}
                  </div>
                  <div className="col-status">
                    {result.error ? (
                      <span className="status-error">超时</span>
                    ) : (
                      <span className="status-success">成功</span>
                    )}
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default TraceRouter
