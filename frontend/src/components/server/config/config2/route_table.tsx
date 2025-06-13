import React from 'react';
import './route_table.css';

const RouteTable: React.FC = () => {
  return (
    <div className="card route-card">
      <h3>Routing Table</h3>
      <table>
        <thead>
          <tr>
            <th>Destination</th>
            <th>Gateway</th>
            <th>Interface</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>0.0.0.0/0</td>
            <td>192.168.0.1</td>
            <td>eth0</td>
          </tr>
          <tr>
            <td>192.168.0.0/24</td>
            <td>0.0.0.0</td>
            <td>eth0</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default RouteTable;
