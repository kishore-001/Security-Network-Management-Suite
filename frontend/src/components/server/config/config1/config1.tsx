import React from 'react';
import './config1.css';
import { useNavigate } from 'react-router-dom';
import ServerConfiguration from './serverconfiguration';
import CommandExecution from './commandexecution';
import UserManagement from './usermanagement';
import SecurityManagement from './securitymanagement';
import ServerOverview from './serveroverview';

const Config1: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div className="config1-container">
      {/* Top Tabs */}
      <div className="top-bar">
        <div className="feature-tabs">
          <button
            className="feature-tab active"
            onClick={() => navigate('/config1')}
          >
            <span className="tab-icon">⚙️</span> General Features
          </button>
          <button
            className="feature-tab"
            onClick={() => navigate('/config2')}
          >
            <span className="tab-icon">🛡️</span> Advanced Features
          </button>
        </div>
      </div>

      {/* Title & Description */}
      <div className="config1-header">
        <div className="config-icon" />
        <h1>Server Configuration Management</h1>
        <p>
          Comprehensive server administration, user management, and security configuration in one unified platform
        </p>
      </div>

      {/* Four Horizontal Cards */}
      <div className="card-row-horizontal">
        <ServerConfiguration />
        <CommandExecution />
        <UserManagement />
        <SecurityManagement />
      </div>

      {/* Overview */}
      <div className="overview-wrapper">
        <ServerOverview />
      </div>
    </>
  );
};

export default Config1;
