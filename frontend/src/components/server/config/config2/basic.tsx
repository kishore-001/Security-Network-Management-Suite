import React, { useState } from 'react';
import './basic.css';
import { FaSave } from 'react-icons/fa';

const Basic: React.FC = () => {
  const [method, setMethod] = useState('Static');
  const [gateway, setGateway] = useState('192.168.1.1'); 
  const [subnet, setSubnet] = useState('255.255.255.0');
  const [dns, setDns] = useState('8.8.8.8');

  const handleChange = () => {
    console.log({ method, gateway, subnet, dns });
  };

  return (
    <div className="config2-basic-card">
      <h2 className="config2-basic-title">Basic</h2>

      <div className="config2-basic-field">
        <label>Method :</label>
        <div className="config2-basic-radio">
          <label>
            <input
              type="radio"
              value="Static"
              checked={method === 'Static'}
              onChange={() => setMethod('Static')}
            />
            Static
          </label>
          <label>
            <input
              type="radio"
              value="DHCP"
              checked={method === 'DHCP'}
              onChange={() => setMethod('DHCP')}
            />
            DHCP
          </label>
        </div>
      </div>

      <div className="config2-basic-field">
        <label>Gateway :</label>
        <input
          type="text"
          value={gateway}
          onChange={(e) => setGateway(e.target.value)}
        />
      </div>

      <div className="config2-basic-field">
        <label>Subnet :</label>
        <input
          type="text"
          value={subnet}
          onChange={(e) => setSubnet(e.target.value)}
        />
      </div>

      <div className="config2-basic-field">
        <label>DNS Server :</label>
        <input
          type="text"
          value={dns}
          onChange={(e) => setDns(e.target.value)}
        />
      </div>

      <button className="config2-basic-button" onClick={handleChange}>
        <FaSave className="config2-basic-icon" />
        Change
      </button>
    </div>
  );
};

export default Basic;
