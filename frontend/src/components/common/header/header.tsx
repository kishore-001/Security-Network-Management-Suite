import React, { useState } from 'react';
import './Header.css';
import { FaBell, FaNetworkWired, FaServer } from 'react-icons/fa';
import { IoIosArrowDown } from 'react-icons/io';
import { useNavigate } from 'react-router-dom';

const Header: React.FC = () => {

  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<'network' | 'server'>('network');
  const [selectedServer, setSelectedServer] = useState('Server 1');
  const [dropdownOpen, setDropdownOpen] = useState(false);


  const servers = ['Server 1', 'Server 2', 'Server 3', 'Server 4'];

  const toggleDropdown = () => setDropdownOpen((prev) => !prev);

  const handleServerSelect = (server: string) => {
    setSelectedServer(server);
    setDropdownOpen(false);
  };

  const servers = ['Server A', 'Server B', 'Server C'];

  return (

    <div className="header-container">
      <div className="header-left">
        <div className="logo-circle">
          <span className="logo-letter">S</span>
          <span className="status-indicator" />
        </div>
        <div>
          <div className="brand-title">SNSMS</div>
          <div className="brand-subtitle">Network Management Suite</div>
        </div>
      </div>

      <div className="header-right">
        <div className="server-dropdown-wrapper">
          <div className="server-dropdown" onClick={toggleDropdown}>
            {selectedServer}
            <IoIosArrowDown className="dropdown-icon" />
          </div>
          {dropdownOpen && (
            <div className="server-dropdown-menu">
              {servers.map((server) => (
                <div
                  key={server}
                  className="server-dropdown-item"
                  onClick={() => handleServerSelect(server)}
                >
                  {server}
                </div>
              ))}
            </div>
          )}
        </div>

        <div className="alert-icon" onClick={()=> navigate("/alert")}>
          <FaBell className="bell-icon" />
                <span style={{ paddingLeft: '5px' }}>Alerts</span>
                <span className="alert-count">3</span>
        </div>

        <div className="toggle-switch">
          <div
            className={`toggle-option ${activeTab === 'network' ? 'active-network' : ''}`}
            onClick={() => setActiveTab('network')}
          >
            <FaNetworkWired />
            <span>Network</span>
          </div>
          <div
            className={`toggle-option ${activeTab === 'server' ? 'active-server' : ''}`}
            onClick={() => setActiveTab('server')}
          >
            <FaServer />
            <span>Server</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;
