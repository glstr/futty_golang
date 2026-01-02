// API配置文件
export const apiConfig = {
  // 后端服务器地址
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8882',
  
  // 聊天接口路径
  chatEndpoint: '/snow/chat',
  
  // 聊天服务类型（默认使用gemini_service）
  defaultChatService: 'gemini_service',
  
  // 获取完整的聊天接口URL
  getChatURL: () => {
    return `${apiConfig.baseURL}${apiConfig.chatEndpoint}`
  }
}

