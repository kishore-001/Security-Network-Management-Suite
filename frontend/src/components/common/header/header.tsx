import React, { useState } from 'react';
import { FaBell, FaNetworkWired, FaServer } from 'react-icons/fa';
import './header.css';

const Header: React.FC = () => {
  const [toggleState, setToggleState] = useState<'network' | 'server'>('network');
  const [selectedServer, setSelectedServer] = useState('Server A');

  const handleToggle = () => {
    setToggleState(prev => (prev === 'network' ? 'server' : 'network'));
  };

  const servers = ['Server A', 'Server B', 'Server C'];

  return (
    <header className="header">
      <div className="header-left">
        <h1 className="logo">SNSMS</h1>
        <p className="title">Network Management Suite</p>
      </div>

      <div className="header-right">
        <select
          className="server-dropdown"
          value={selectedServer}
          onChange={(e) => setSelectedServer(e.target.value)}
        >
          {servers.map(server => (
            <option key={server} value={server}>{server}</option>
          ))}
        </select>

        <div className={`custom-toggle ${toggleState}`} onClick={handleToggle}>
          <div className="toggle-slider">
            {toggleState === 'network' ? <FaNetworkWired /> : <FaServer />}
          </div>
          <div className="toggle-labels">
            <span>Network</span>
            <span>Server</span>
          </div>
        </div>

        <button className="action-btn alerts">
          <FaBell className="icon" />
          <span>Alerts</span>
          <span className="alert-badge">3</span>
        </button>
      </div>
    </header>
  );
};

export default Header;
