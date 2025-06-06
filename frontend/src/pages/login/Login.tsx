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

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await response.json();

      if (response.ok && data.status === 'ok') {
        // Save token to localStorage
        localStorage.setItem('token', data.token);
        setSuccess(true);

        // Optional: Redirect after short delay
        setTimeout(() => {
          window.location.href = '/dashboard'; // or use navigate() if using React Router
        }, 1000);
      } else if (response.status === 401) {
        setError('Invalid username or password.');
      } else {
        setError('Something went wrong. Please try again.');
      }
    } catch (err) {
      console.error(err);
      setError('Network error. Check your connection.');
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
                aria-label="Username"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                autoFocus
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
                aria-label="Password"
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
              <a href="#" onClick={(e) => e.preventDefault()}>
                Forgot password?
              </a>
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
