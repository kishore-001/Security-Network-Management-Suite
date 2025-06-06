import React, { useState } from 'react';
import './login.css';
import { FaUser, FaLock } from 'react-icons/fa';

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);

  const handleLogin = async (): Promise<void> => {
    setError(null);
    setSuccess(false);
    setLoading(true);

    // Simulate a backend delay
    await new Promise((res) => setTimeout(res, 1000));

    // Mock validation logic
    if (!username || !password) {
      setError('Username and password are required.');
    } else if (username === 'admin' && password === 'admin123') {
      // Successful login case
      setSuccess(true);
    } else {
      // Failed login case
      setError('Invalid username or password.');
    }

    setLoading(false);
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

          <form
            onSubmit={(e) => {
              e.preventDefault();
              handleLogin();
            }}
          >
            <label htmlFor="username">Username</label>
            <div className="input-group">
              <span className="icon">
                <FaUser />
              </span>
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
              <span className="icon">
                <FaLock />
              </span>
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

            <button type="submit" disabled={loading}>
              {loading ? 'Signing In...' : 'Sign In'}
            </button>
          </form>

          {/* Show success or error messages */}
          {error && <p style={{ color: 'red', marginTop: '10px' }}>{error}</p>}
          {success && <p style={{ color: 'green', marginTop: '10px' }}>Login successful! 🎉</p>}

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
