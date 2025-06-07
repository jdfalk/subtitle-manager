import { useState } from 'react'
import './App.css'

function App() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [status, setStatus] = useState('')

  const login = async () => {
    const form = new URLSearchParams({ username, password })
    const res = await fetch('/api/login', { method: 'POST', body: form })
    if (res.ok) {
      setStatus('logged in')
    } else {
      setStatus('login failed')
    }
  }

  return (
    <div className="login">
      <h1>Subtitle Manager</h1>
      <input placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
      <button onClick={login}>Login</button>
      <p>{status}</p>
    </div>
  )
}

export default App
