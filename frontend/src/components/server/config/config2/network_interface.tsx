import React, { useState } from 'react';
import './network_interface.css';
import { FiRefreshCw } from 'react-icons/fi';

const interfaces = ['eth0', 'eth1', 'wlan0', 'wlan1'];

const NetworkInterface: React.FC = () => {
  const [status, setStatus] = useState<Record<string, string>>({
    eth0: '',
    eth1: 'enabled',
    wlan0: 'enabled',
    wlan1: 'enabled',
  });

  const [restarting, setRestarting] = useState(false);

  const handleToggle = (iface: string, action: 'enable' | 'disable') => {
    setStatus((prev) => ({
      ...prev,
      [iface]: action === 'enable' ? 'enabled' : 'disabled',
    }));
  };

  const handleRestart = () => {
    setRestarting(true);
    setTimeout(() => {
      setRestarting(false);
    }, 1500);
  };

  return (
    <div className="card config2-network-card">
      <h3>Network Interface</h3>
      <p className="config2-network-subtitle">Interface :</p>

      <div className="config2-interface-list">
        {interfaces.map((iface) => (
          <div className="config2-interface-row" key={iface}>
            <span>{iface}</span>
            <div className="config2-buttons">
              <button
                className={`config2-btn ${
                  status[iface] === 'enabled' ? 'enabled' : ''
                }`}
                onClick={() => handleToggle(iface, 'enable')}
              >
                Enable
              </button>
              <button
                className={`config2-btn ${
                  status[iface] === 'disabled' ? 'disabled' : ''
                }`}
                onClick={() => handleToggle(iface, 'disable')}
              >
                Disable
              </button>
            </div>
          </div>
        ))}
      </div>

      <div className="config2-status-uptime">
        <div>
          <p>Status :</p>
          <div className="config2-box green">Online</div>
        </div>
        <div>
          <p>Uptime :</p>
          <div className="config2-box">19:00:00</div>
        </div>
      </div>

      <button
        className={`config2-restart-btn ${restarting ? 'spinning' : ''}`}
        onClick={handleRestart}
      >
        <FiRefreshCw className="icon" />
        Restart Service
      </button>
    </div>
  );
};

export default NetworkInterface;
