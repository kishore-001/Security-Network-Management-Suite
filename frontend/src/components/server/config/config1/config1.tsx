import React from 'react';
import './config1.css';
import ServerConfiguration from './serverconfiguration';
import CommandExecution from './commandexecution';
 import UserManagement from './usermanagement';
 import SecurityManagement from './securitymanagement';
import ServerOverview from './serveroverview';

const Config1: React.FC = () => {
  return (
    <div className="config1-container">
      <div className="config1-header">
        <div className="icon" />
        <h1>Server Configuration Management</h1>
        <p>
          Comprehensive server administration, user management, and security configuration in one unified platform
        </p>
      </div>
      <div className="card-row">
        
       <ServerConfiguration />
        <CommandExecution />
        { <UserManagement /> }
        { <SecurityManagement /> }
      </div>
      <ServerOverview />
    </div>
  );
};

export default Config1;