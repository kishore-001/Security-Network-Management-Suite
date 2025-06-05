  import React, { useState } from 'react';
  import './login.css';
  import { FaUser, FaLock } from 'react-icons/fa';

  const Login: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

 const handleLogin = async (): Promise<void> => {
  console.log('Attempting login with:', { username, password });

  try {
    await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });
  } catch (error) {
    console.error('Login error:', error);
  }
};



    return (
      <>
        <div className="logo-outside">
          <FaLock />
        </div>

        <div className="login-container">
          <div className="login-card">
            <h1>SNSMS</h1>
            <p className="subtext">Network Management Suite</p>

            <form onSubmit={(e) => { e.preventDefault(); handleLogin(); }}>
              <label htmlFor="username">Username</label>
              <div className="input-group">
                <span className="icon"><FaUser /></span>
                <input
                  type="text"
                  id="username"
                  placeholder="Username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </div>

              <label htmlFor="password">Password</label>
              <div className="input-group">
                <span className="icon"><FaLock /></span>
                <input
                  type="password"
                  id="password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>

              <div className="options-row">
                <label>
                  <input type="checkbox" style={{ marginRight: '6px' }} />
                  Remember me
                </label>
                <a href="#">Forgot password?</a>
              </div>

              <button type="submit">Sign In</button>
            </form>

            <div className="footer">
              Don’t have an account? <a href="#">Contact administrator</a>
              <br />
              <span style={{ fontSize: '0.75rem', marginTop: '6px', display: 'inline-block' }}>
                © 2025 SNSMS Network Management Suite. All rights reserved.
              </span>
            </div>
          </div>
        </div>
      </>
    );
  };

  export default Login;
