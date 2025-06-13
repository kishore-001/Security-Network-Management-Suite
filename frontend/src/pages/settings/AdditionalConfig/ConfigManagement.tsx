import React from 'react';
import './AdditionalConfig.css';

const ConfigManagement: React.FC = () => {
  return (
    <div className="config-management-container">
      <div className="action-buttons">
        <button className="btn update">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
            <path d="M21 15V19C21 20.1046 20.1046 21 19 21H5C3.89543 21 3 20.1046 3 19V5C3 3.89543 3.89543 3 5 3H9" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            <path d="M16 3H21V8" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            <path d="M11 13L21 3" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
          </svg>
          SNSMS Config Download
        </button>
        <button className="btn restore">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
            <path d="M21 12V19C21 20.1046 20.1046 21 19 21H5C3.89543 21 3 20.1046 3 19V5C3 3.89543 3.89543 3 5 3H12" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            <path d="M18 3V8H21" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            <path d="M11 13L3 21" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
            <path d="M11 21H3V13" stroke="currentColor" strokeWidth="2" strokeLinecap="round"/>
          </svg>
          SNSMS Restore Config
        </button>
      </div>
    </div>
  );
};

export default ConfigManagement;