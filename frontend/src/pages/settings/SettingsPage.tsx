import React, { useState } from 'react';
import Sidebar from '../../components/common/sidebar/sidebar';
import Header from '../../components/common/header/header';
import SettingsTabs from './SettingsTabs';
import UserManagement from './UserManagement/UserManagement';
import AdditionalConfig from './AdditionalConfig/AdditionalConfig';
import './settings.css';

const SettingsPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<'user' | 'config'>('user');

  return (
    <div className="page-layout">
      {/* HEADER */}
      <Header />

      {/* MAIN SECTION */}
      <div className="main-section">
        {/* Sidebar now renders Health internally */}
        <div className="left-panel">
          <Sidebar />
        </div>

        {/* Right-hand content */}
        <div className="content-container">
          <SettingsTabs activeTab={activeTab} setActiveTab={setActiveTab} />
          {activeTab === 'user' ? <UserManagement /> : <AdditionalConfig />}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;
