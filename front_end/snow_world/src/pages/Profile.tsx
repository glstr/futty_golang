import './Profile.css'

function Profile() {
  return (
    <div className="profile-container">
      <div className="profile-header">
        <div className="profile-avatar">
          <div className="avatar-circle">
            <span className="avatar-text">我</span>
          </div>
        </div>
        <h1 className="profile-name">我的个人主页</h1>
        <p className="profile-bio">欢迎来到我的个人空间</p>
      </div>

      <div className="profile-content">
        <section className="profile-section">
          <h2 className="section-title">关于我</h2>
          <div className="section-content">
            <p>这里可以介绍自己的基本信息、兴趣爱好、专业技能等。</p>
            <p>你可以随时修改这个页面来展示你想要的内容。</p>
          </div>
        </section>

        <section className="profile-section">
          <h2 className="section-title">技能</h2>
          <div className="section-content">
            <div className="skills-list">
              <span className="skill-tag">React</span>
              <span className="skill-tag">TypeScript</span>
              <span className="skill-tag">Go</span>
              <span className="skill-tag">Vite</span>
            </div>
          </div>
        </section>

        <section className="profile-section">
          <h2 className="section-title">联系方式</h2>
          <div className="section-content">
            <p>📧 Email: your.email@example.com</p>
            <p>🔗 GitHub: github.com/yourusername</p>
          </div>
        </section>
      </div>
    </div>
  )
}

export default Profile

