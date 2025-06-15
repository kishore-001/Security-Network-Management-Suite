// components/server/alert/AlertDashboard.tsx
import React, { useState } from 'react';
import { FaSync, FaDownload, FaExclamationTriangle, FaChartLine, FaFilter, FaClock, FaCog } from 'react-icons/fa';
import './AlertDashboard.css';

const AlertDashboard: React.FC = () => {
  const [activeTab, setActiveTab] = useState<'alerts' | 'performance'>('alerts');
  const [selectedFilter, setSelectedFilter] = useState('All Alerts');

  // Mock data - will be replaced with real API data
  const issuesData = {
    total: 11,
    new: 10,
    inProgress: 0,
    resolved: 1
  };

  const alertsData = {
    total: 1551,
    high: 1401,
    medium: 39,
    low: 111
  };

  const activeAlerts = [
    {
      id: 1,
      severity: 'high',
      message: 'CPU usage exceeded 90% on server-01',
      time: '10 minutes ago',
      icon: 'cpu'
    },
    {
      id: 2,
      severity: 'high',
      message: 'Disk space on /var is 98% full',
      time: '15 minutes ago',
      icon: 'disk'
    },
    {
      id: 3,
      severity: 'high',
      message: 'Service nginx crashed unexpectedly',
      time: '20 minutes ago',
      icon: 'service'
    },
    {
      id: 4,
      severity: 'high',
      message: 'High number of failed SSH login attempts',
      time: '25 minutes ago',
      icon: 'security'
    },
    {
      id: 5,
      severity: 'high',
      message: 'Server unreachable',
      time: '30 minutes ago',
      icon: 'network'
    },
    {
      id: 6,
      severity: 'medium',
      message: 'Disk I/O anomaly',
      time: '35 minutes ago',
      icon: 'disk'
    }
  ];

  const getSeverityIcon = (severity: string) => {
    const iconMap = {
      cpu: '💻',
      disk: '💾',
      service: '⚙️',
      security: '🔒',
      network: '🌐'
    };
    return iconMap[severity as keyof typeof iconMap] || '⚠️';
  };

  const getSeverityClass = (severity: string) => {
    return `monitoring-alerts-severity-${severity}`;
  };

  return (
    <div className="monitoring-alerts-dashboard">
      {/* Header Section */}
      <div className="monitoring-alerts-header-section">
        <div className="monitoring-alerts-title-section">
          <h1 className="monitoring-alerts-page-title">Monitoring Dashboard</h1>
          <p className="monitoring-alerts-page-subtitle">Real-time system monitoring and alert management</p>
        </div>
        <div className="monitoring-alerts-header-actions">
          <button className="monitoring-alerts-btn monitoring-alerts-btn-secondary">
            <FaSync className="monitoring-alerts-btn-icon" />
            Refresh Data
          </button>
          <button className="monitoring-alerts-btn monitoring-alerts-btn-primary">
            <FaDownload className="monitoring-alerts-btn-icon" />
            Export Report
          </button>
        </div>
      </div>

      {/* Stats Section */}
      <div className="monitoring-alerts-stats-container">
        {/* Issues Stats */}
        <div className="monitoring-alerts-stats-card">
          <div className="monitoring-alerts-stats-header">
            <h3>Issues</h3>
          </div>
          <div className="monitoring-alerts-stats-grid">
            <div className="monitoring-alerts-stat-item">
              <span className="monitoring-alerts-stat-value">{issuesData.total}</span>
              <span className="monitoring-alerts-stat-label">Total</span>
            </div>
            <div className="monitoring-alerts-stat-item monitoring-alerts-stat-new">
              <span className="monitoring-alerts-stat-value">{issuesData.new}</span>
              <span className="monitoring-alerts-stat-label">New</span>
            </div>
            <div className="monitoring-alerts-stat-item monitoring-alerts-stat-progress">
              <span className="monitoring-alerts-stat-value">{issuesData.inProgress}</span>
              <span className="monitoring-alerts-stat-label">In Progress</span>
            </div>
            <div className="monitoring-alerts-stat-item monitoring-alerts-stat-resolved">
              <span className="monitoring-alerts-stat-value">{issuesData.resolved}</span>
              <span className="monitoring-alerts-stat-label">Resolved</span>
            </div>
          </div>
          <div className="monitoring-alerts-progress-bar">
            <div className="monitoring-alerts-progress-fill" style={{ width: '9.1%' }}></div>
          </div>
          <span className="monitoring-alerts-progress-text">9.1%</span>
        </div>

        {/* Alerts Stats */}
        <div className="monitoring-alerts-stats-card">
          <div className="monitoring-alerts-stats-header">
            <h3>Alerts</h3>
          </div>
          <div className="monitoring-alerts-donut-container">
            <div className="monitoring-alerts-donut-chart">
              <div className="monitoring-alerts-donut-center">
                <span className="monitoring-alerts-donut-total">{alertsData.total}</span>
                <span className="monitoring-alerts-donut-label">Total</span>
              </div>
            </div>
            <div className="monitoring-alerts-donut-stats">
              <div className="monitoring-alerts-donut-stat monitoring-alerts-high">
                <span className="monitoring-alerts-donut-value">{alertsData.high}</span>
                <span className="monitoring-alerts-donut-stat-label">High</span>
              </div>
              <div className="monitoring-alerts-donut-stat monitoring-alerts-medium">
                <span className="monitoring-alerts-donut-value">{alertsData.medium}</span>
                <span className="monitoring-alerts-donut-stat-label">Medium</span>
              </div>
              <div className="monitoring-alerts-donut-stat monitoring-alerts-low">
                <span className="monitoring-alerts-donut-value">{alertsData.low}</span>
                <span className="monitoring-alerts-donut-stat-label">Low</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Tabs Section */}
      <div className="monitoring-alerts-tabs-container">
        <div className="monitoring-alerts-tabs">
          <button 
            className={`monitoring-alerts-tab ${activeTab === 'alerts' ? 'active' : ''}`}
            onClick={() => setActiveTab('alerts')}
          >
            <FaExclamationTriangle className="monitoring-alerts-tab-icon" />
            Alerts
          </button>
          <button 
            className={`monitoring-alerts-tab ${activeTab === 'performance' ? 'active' : ''}`}
            onClick={() => setActiveTab('performance')}
          >
            <FaChartLine className="monitoring-alerts-tab-icon" />
            Performance
          </button>
        </div>
        <div className="monitoring-alerts-filter-section">
          <select 
            className="monitoring-alerts-filter-select"
            value={selectedFilter}
            onChange={(e) => setSelectedFilter(e.target.value)}
          >
            <option value="All Alerts">All Alerts</option>
            <option value="High Priority">High Priority</option>
            <option value="Medium Priority">Medium Priority</option>
            <option value="Low Priority">Low Priority</option>
          </select>
          <FaFilter className="monitoring-alerts-filter-icon" />
        </div>
      </div>

      {/* Active Alerts Section */}
      <div className="monitoring-alerts-content-container">
        <div className="monitoring-alerts-section-header">
          <FaExclamationTriangle className="monitoring-alerts-section-icon" />
          <h3>Active Alerts</h3>
        </div>

        <div className="monitoring-alerts-table-container">
          <table className="monitoring-alerts-table">
            <thead>
              <tr>
                <th>Severity</th>
                <th>Alert</th>
                <th>Time</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {activeAlerts.map((alert) => (
                <tr key={alert.id} className="monitoring-alerts-table-row">
                  <td>
                    <span className={`monitoring-alerts-severity-badge ${getSeverityClass(alert.severity)}`}>
                      {alert.severity}
                    </span>
                  </td>
                  <td>
                    <div className="monitoring-alerts-message">
                      <span className="monitoring-alerts-icon">{getSeverityIcon(alert.icon)}</span>
                      <span className="monitoring-alerts-text">{alert.message}</span>
                    </div>
                  </td>
                  <td>
                    <div className="monitoring-alerts-time">
                      <FaClock className="monitoring-alerts-time-icon" />
                      <span>{alert.time}</span>
                    </div>
                  </td>
                  <td>
                    <button className="monitoring-alerts-action-btn">
                      <FaCog />
                      Resolve
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default AlertDashboard;
