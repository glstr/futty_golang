import './ESStorage.css'
import { useMemo } from 'react'
import { apiConfig } from '../config/api'

function ESStorage() {
  const kibanaURL = useMemo(() => {
    const base = (apiConfig as any).kibanaBaseURL || 'http://localhost:5601'
    return `${base}/app`
  }, [])

  return (
    <div className="es-storage-container">
      <div className="es-storage-header">
        <h1>ES存储</h1>
        <p className="es-storage-subtitle">内嵌 Kibana 页面，支持页面内跳转</p>
      </div>
      <div className="es-storage-frame-wrapper">
        <iframe
          title="Kibana"
          src={kibanaURL}
          className="es-storage-iframe"
        />
      </div>
    </div>
  )
}

export default ESStorage
