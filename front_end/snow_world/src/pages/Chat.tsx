import { useState, useRef, useEffect } from 'react'
import './Chat.css'
import { apiConfig } from '../config/api'

interface Message {
  id: number
  text: string
  sender: 'user' | 'bot'
  timestamp: Date
}

function Chat() {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: 1,
      text: '你好！欢迎使用聊天功能。',
      sender: 'bot',
      timestamp: new Date()
    }
  ])
  const [inputValue, setInputValue] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const messagesEndRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLTextAreaElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const handleSend = async () => {
    if (inputValue.trim() === '' || isLoading) return

    const userMessageText = inputValue.trim()
    const newMessage: Message = {
      id: messages.length + 1,
      text: userMessageText,
      sender: 'user',
      timestamp: new Date()
    }

    setMessages([...messages, newMessage])
    setInputValue('')
    setIsLoading(true)
    
    // 重置输入框高度
    if (inputRef.current) {
      inputRef.current.style.height = 'auto'
    }

    try {
      // 调用后端接口
      const response = await fetch(apiConfig.getChatURL(), {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          service: apiConfig.defaultChatService,
          message: userMessageText,
        }),
      })

      const data = await response.json()
      
      if (!response.ok || data.error_code !== 0) {
        throw new Error(data.error_msg || '请求失败')
      }

      const botMessage: Message = {
        id: messages.length + 2,
        text: data.response || '抱歉，我暂时无法回复。',
        sender: 'bot',
        timestamp: new Date()
      }
      setMessages(prev => [...prev, botMessage])
    } catch (error) {
      console.error('Chat API error:', error)
      const errorMessage: Message = {
        id: messages.length + 2,
        text: '抱歉，发生错误：' + (error instanceof Error ? error.message : '未知错误'),
        sender: 'bot',
        timestamp: new Date()
      }
      setMessages(prev => [...prev, errorMessage])
    } finally {
      setIsLoading(false)
      // 聚焦输入框
      inputRef.current?.focus()
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      handleSend()
    }
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputValue(e.target.value)
    // 自动调整输入框高度
    if (inputRef.current) {
      inputRef.current.style.height = 'auto'
      inputRef.current.style.height = `${Math.min(inputRef.current.scrollHeight, 150)}px`
    }
  }

  const formatTime = (date: Date) => {
    return date.toLocaleTimeString('zh-CN', { 
      hour: '2-digit', 
      minute: '2-digit' 
    })
  }

  return (
    <div className="chat-container">
      <div className="chat-header">
        <h1>聊天</h1>
      </div>
      
      <div className="chat-messages">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`message ${message.sender === 'user' ? 'message-user' : 'message-bot'}`}
          >
            <div className="message-content">
              <div className="message-text">{message.text}</div>
              <div className="message-time">{formatTime(message.timestamp)}</div>
            </div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>

      <div className="chat-input-container">
        <div className="chat-input-wrapper">
          <textarea
            ref={inputRef}
            className="chat-input"
            value={inputValue}
            onChange={handleInputChange}
            onKeyDown={handleKeyDown}
            placeholder="输入消息... (按 Enter 发送，Shift+Enter 换行)"
            rows={1}
          />
          <button
            className="chat-send-button"
            onClick={handleSend}
            disabled={inputValue.trim() === '' || isLoading}
          >
            {isLoading ? '发送中...' : '发送'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default Chat

