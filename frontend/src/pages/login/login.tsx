import React from 'react';
import './login.css';
import { FaUser, FaLock } from 'react-icons/fa';

const Login: React.FC = () => {
  return (
    <div className="login-container">
      <div className="login-card">
        <div className="logo-icon">
          <FaLock />
        </div>
        <h1>SNSMS</h1>
        <p className="subtext">Network Management Suite</p>

        <form>
          <label htmlFor="username">Username</label>
          <div className="input-group">
            <span className="icon"><FaUser /></span>
            <input type="text" id="username" placeholder="Username" />
          </div>

          <label htmlFor="password">Password</label>
          <div className="input-group">
            <span className="icon"><FaLock /></span>
            <input type="password" id="password" placeholder="Password" />
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
  );
};

export default Login;
