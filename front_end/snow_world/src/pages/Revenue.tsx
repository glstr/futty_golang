import './Revenue.css'

interface DailyRevenue {
  date: string
  value: number
}

const sampleData: DailyRevenue[] = [
  { date: '2025-01-01', value: 1200 },
  { date: '2025-01-02', value: 1500 },
  { date: '2025-01-03', value: 980 },
  { date: '2025-01-04', value: 1850 },
  { date: '2025-01-05', value: 2100 },
  { date: '2025-01-06', value: 1700 },
  { date: '2025-01-07', value: 2400 },
  { date: '2025-01-08', value: 2600 },
  { date: '2025-01-09', value: 1950 },
  { date: '2025-01-10', value: 2800 }
]

function Revenue() {
  const data = sampleData

  const width = 900
  const height = 360
  const marginTop = 40
  const marginRight = 40
  const marginBottom = 80
  const marginLeft = 80

  const values = data.map((d) => d.value)
  const maxValue = Math.max(...values, 0)
  const minValue = 0

  const innerWidth = width - marginLeft - marginRight
  const innerHeight = height - marginTop - marginBottom

  const xStep = data.length > 1 ? innerWidth / (data.length - 1) : 0

  const getX = (index: number) => marginLeft + xStep * index
  const getY = (value: number) => {
    if (maxValue === minValue) {
      return marginTop + innerHeight / 2
    }
    const ratio = (value - minValue) / (maxValue - minValue)
    return marginTop + innerHeight * (1 - ratio)
  }

  const points = data
    .map((d, index) => `${getX(index)},${getY(d.value)}`)
    .join(' ')

  const yTicks = 4
  const yTickValues = Array.from({ length: yTicks + 1 }, (_, i) => {
    if (maxValue === 0) {
      return 0
    }
    return Math.round((maxValue / yTicks) * i)
  })

  const total = values.reduce((sum, v) => sum + v, 0)
  const average = data.length > 0 ? total / data.length : 0
  const last = data.length > 0 ? data[data.length - 1].value : 0

  return (
    <div className="revenue-container">
      <div className="revenue-header">
        <h1>收益数据展示</h1>
        <p className="revenue-subtitle">
          横轴为日期（天），纵轴为收益（元）
        </p>
      </div>

      <div className="revenue-summary">
        <div className="summary-item">
          <div className="summary-label">总收益</div>
          <div className="summary-value">{total.toFixed(2)} 元</div>
        </div>
        <div className="summary-item">
          <div className="summary-label">日均收益</div>
          <div className="summary-value">{average.toFixed(2)} 元</div>
        </div>
        <div className="summary-item">
          <div className="summary-label">最近一天收益</div>
          <div className="summary-value">{last.toFixed(2)} 元</div>
        </div>
      </div>

      <div className="revenue-chart-card">
        <div className="revenue-chart-header">
          <h2>收益曲线</h2>
          <span className="revenue-chart-unit">单位：元</span>
        </div>
        <div className="revenue-chart-wrapper">
          <svg
            viewBox={`0 0 ${width} ${height}`}
            className="revenue-chart-svg"
          >
            <line
              x1={marginLeft}
              y1={marginTop}
              x2={marginLeft}
              y2={height - marginBottom}
              className="axis-line"
            />
            <line
              x1={marginLeft}
              y1={height - marginBottom}
              x2={width - marginRight}
              y2={height - marginBottom}
              className="axis-line"
            />

            {yTickValues.map((v, i) => {
              const y = getY(v)
              return (
                <g key={i}>
                  <line
                    x1={marginLeft}
                    y1={y}
                    x2={width - marginRight}
                    y2={y}
                    className="grid-line"
                  />
                  <text
                    x={marginLeft - 10}
                    y={y}
                    className="axis-label axis-label-y"
                  >
                    {v.toFixed(0)}
                  </text>
                </g>
              )
            })}

            {data.map((d, index) => {
              const showLabel =
                data.length <= 10 || index === 0 || index === data.length - 1 || index % 2 === 0
              if (!showLabel) {
                return null
              }
              const x = getX(index)
              const y = height - marginBottom
              return (
                <g key={d.date}>
                  <line
                    x1={x}
                    y1={y}
                    x2={x}
                    y2={y + 6}
                    className="tick-line"
                  />
                  <text
                    x={x}
                    y={y + 20}
                    className="axis-label axis-label-x"
                  >
                    {d.date.slice(5)}
                  </text>
                </g>
              )
            })}

            {data.length > 0 && (
              <>
                <polyline
                  points={points}
                  className="revenue-line"
                />
                {data.map((d, index) => {
                  const x = getX(index)
                  const y = getY(d.value)
                  return (
                    <g key={d.date}>
                      <circle
                        cx={x}
                        cy={y}
                        r={4}
                        className="revenue-point"
                      />
                    </g>
                  )
                })}
              </>
            )}
          </svg>
        </div>
      </div>
    </div>
  )
}

export default Revenue

