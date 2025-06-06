import React, { useState } from 'react';
import { FaBell, FaNetworkWired, FaServer } from 'react-icons/fa';
import './header.css';

const TopNav: React.FC = () => {
  const [toggleState, setToggleState] = useState<'network' | 'server'>('network');

  const handleToggle = () => {
    setToggleState(prev => (prev === 'network' ? 'server' : 'network'));
  };

  return (
    <header className="snsms-header">
      {/* Left: Logo and Title */}
      <div className="header-brand">
        <h1 className="logo">SNSMS</h1>
        <p className="title">Network Management Suite</p>
      </div>

      {/* Center: Fancy Toggle */}
      <div className="toggle-container">
        <div className={`toggle-wrapper ${toggleState}`} onClick={handleToggle}>
          <div className="toggle-option network">
            <FaNetworkWired className="toggle-icon" />
            <span>Network</span>
          </div>
          <div className="toggle-option server">
            <FaServer className="toggle-icon" />
            <span>Server</span>
          </div>
          <div className="toggle-slider" />
        </div>
      </div>

      {/* Right: Alerts */}
      <div className="header-actions">
        <button className="action-btn alerts">
          <FaBell className="icon" />
          <span>Alerts</span>
          <span className="alert-badge">3</span>
        </button>
      </div>
    </header>
  );
};

export default TopNav;
