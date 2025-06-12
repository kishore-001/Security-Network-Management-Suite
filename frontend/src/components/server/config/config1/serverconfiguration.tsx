import React, { useState } from 'react';
import './serverconfiguration.css';
import { FaServer } from 'react-icons/fa';
import ModalWrapper from './modalwrapper';

const ServerConfiguration: React.FC = () => {
  const [showModal, setShowModal] = useState(false);

  const handleApply = (hostname: string, timezone: string) => {
    console.log("Applied:", hostname, timezone);
    setShowModal(false);
  };

  return (
    <>
      <div className="card server-card" onClick={() => setShowModal(true)}>
        <div className="server-icon">
          <FaServer size={20} color="white" />
        </div>
        <h3>System Configuration</h3>
        <p>Configure server hostname and timezone settings</p>
      </div>

      {showModal && (
        <ModalWrapper title="System Configuration" onClose={() => setShowModal(false)}>
          <p className="subtitle">Configure fundamental server parameters and system-level settings</p>
          <form onSubmit={(e) => {
            e.preventDefault();
            const hostname = (e.target as any).hostname.value;
            const timezone = (e.target as any).timezone.value;
            handleApply(hostname, timezone);
          }}>
            <label>Server Hostname</label>
            <input name="hostname" placeholder="server.company.com" required />
            
            <label>System Timezone</label>
            <select name="timezone">
              <option value="UTC+0 (GMT)">UTC+0 (GMT)</option>
              <option value="UTC+5:30 (IST)">UTC+5:30 (IST)</option>
              <option value="UTC-8 (PST)">UTC-8 (PST)</option>
            </select>

            <button type="submit"><FaServer /> Apply Configuration</button>
          </form>
        </ModalWrapper>
      )}
    </>
  );
};

export default ServerConfiguration;
