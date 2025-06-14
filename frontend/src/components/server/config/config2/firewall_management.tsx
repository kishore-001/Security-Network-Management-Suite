import React from 'react';
import './firewall_management.css';
import pen from '../../../../assets/icon/pen.svg'; 
import x from '../../../../assets/icon/x.svg'
import eye from '../../../../assets/icon/eye.svg'
const FirewallManagement: React.FC = () => {
  const firewallRules = [
    {
      rule: 'LAN',
      src: 'LAN',
      dest: 'LAN',
      protocol: 'Any',
      srcPort: 'Any',
      destPort: 'Any',
      action: 'Allow',
      status: 'Active',
      log: true,
      schedule: true,
    },
    {
      rule: 'WAN',
      src: 'WAN',
      dest: 'LAN',
      protocol: 'TCP',
      srcPort: 'Any',
      destPort: '80,443',
      action: 'Allow',
      status: 'Active',
      log: false,
      schedule: false,
    },
    {
      rule: 'Block',
      src: 'Any',
      dest: 'WAN',
      protocol: 'ICMP',
      srcPort: 'Any',
      destPort: 'Any',
      action: 'Block',
      status: 'Active',
      log: true,
      schedule: true,
    },
  ];

  return (
    <div className="firewall-config2-card">
      <div className="firewall-config2-header">
        <i className="firewall-config2-icon-shield" /> <h3>Firewall management</h3>
      </div>
      <div className="firewall-config2-table-wrapper">
        <table className="firewall-config2-table">
          <thead>
            <tr>
              <th>Rule</th>
              <th>Src</th>
              <th>Dest</th>
              <th>Protocol</th>
              <th>Src Port</th>
              <th>Dest Port</th>
              <th>Action</th>
              <th>Status</th>
              <th>Log</th>
              <th>Schedule</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {firewallRules.map((rule, index) => (
              <tr key={index}>
                <td>{rule.rule}</td>
                <td>{rule.src}</td>
                <td>{rule.dest}</td>
                <td>{rule.protocol}</td>
                <td>{rule.srcPort}</td>
                <td>{rule.destPort}</td>
                <td>
                  <span className={`firewall-config2-action-button ${rule.action.toLowerCase()}`}>
                    {rule.action}
                  </span>
                </td>
                <td>
                  <span className="firewall-config2-status-active">Active</span>
                </td>
                <td>
                  <span className={`firewall-config2-dot ${rule.log ? 'green' : 'gray'}`}></span>
                </td>
                <td>
                  <span className={`firewall-config2-dot ${rule.schedule ? 'blue' : 'gray'}`}></span>
                </td>
                <td className="firewall-config2-action-icons">
                  <button className="firewall-config2-icon-btn edit" title="Edit"><img src={pen} alt="pen" />
                   </button>
                  <button className="firewall-config2-icon-btn view" title="View"> <img src={eye} alt="eye" />
                   </button>
                  <button className="firewall-config2-icon-btn delete" title="Delete"><img src={x} alt="x" />
                   </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default FirewallManagement;
