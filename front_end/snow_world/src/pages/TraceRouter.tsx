import { useState } from 'react'
import './TraceRouter.css'
import { apiConfig } from '../config/api'

interface TraceResult {
  ttl: number
  network: string
  addr: string
  duration: string
  error?: string
}

function TraceRouter() {
  const [ipAddress, setIpAddress] = useState('')
  const [results, setResults] = useState<TraceResult[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleTrace = async () => {
    if (!ipAddress.trim()) {
      setError('请输入IP地址')
      return
    }

    // 简单的IP地址验证
    const ipRegex = /^(\d{1,3}\.){3}\d{1,3}$/
    if (!ipRegex.test(ipAddress.trim())) {
      setError('请输入有效的IP地址格式（如：8.8.8.8）')
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
          if (attempts < maxAttempts) {
            attempts++
            setTimeout(poll, 1000) // 每秒轮询一次
          } else {
            setError(data.error_msg || '请求超时，请稍后重试')
            setIsLoading(false)
          }
          return
        }

        if (data.result) {
          // 解析result字符串（可能是JSON格式）
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
          setIsLoading(false)
        } else if (attempts < maxAttempts) {
          attempts++
          setTimeout(poll, 1000) // 继续轮询
        } else {
          setError('请求超时，请稍后重试')
          setIsLoading(false)
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
        <p className="trace-router-subtitle">输入IP地址，追踪数据包经过的路由节点</p>
      </div>

      <div className="trace-router-input-section">
        <div className="input-wrapper">
          <input
            type="text"
            className="trace-input"
            value={ipAddress}
            onChange={(e) => setIpAddress(e.target.value)}
            onKeyPress={handleKeyPress}
            placeholder="请输入IP地址（如：8.8.8.8）"
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

