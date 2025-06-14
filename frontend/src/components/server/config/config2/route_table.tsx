import React, { useState } from 'react';
import './route_table.css';
import router from '../../../../assets/icon/router.svg'
import bin from '../../../../assets/icon/trash-2.svg';
interface Route {
  serial: string;
  route: string;
  time: string;
}

const RouteTable: React.FC = () => {
  const [routes, setRoutes] = useState<Route[]>([
    { serial: '003', route: '172.16.0.0/12', time: '14:32:45' },
    { serial: '002', route: 'jdhcskd', time: '08:45:54' }
  ]);

  const [input, setInput] = useState('');
  const [serialCounter, setSerialCounter] = useState(4);

  const handleAdd = () => {
    if (!input.trim()) return;
    const now = new Date().toLocaleTimeString('en-GB');
    const newRoute: Route = {
      serial: `00${serialCounter}`,
      route: input.trim(),
      time: now
    };
    setRoutes([newRoute, ...routes]);
    setSerialCounter(prev => prev + 1);
    setInput('');
  };

  const handleDelete = (index: number) => {
    const updated = routes.filter((_, i) => i !== index);
    setRoutes(updated);
  };

  return (
    <div className="card config2-route-card">
      <div className="config2-route-header">
        <img src= {router} alt="Routing Icon" className="config2-route-icon" />
        <h3>Route Table</h3>
      </div>

      <div className="config2-route-inputs">
        <input
          type="text"
          placeholder="Enter route"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button className="config2-add-btn" onClick={handleAdd}>Add</button>
      </div>

      <table className="config2-route-table">
        <thead>
          <tr>
            <th>SERIAL NO</th>
            <th>ROUTE</th>
            <th>TIME</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {routes.map((route, index) => (
            <tr key={index}>
              <td>{route.serial}</td>
              <td>{route.route}</td>
              <td>{route.time}</td>
              <td>
                <button className="config2-delete-btn" onClick={() => handleDelete(index)}>
                  <img src={bin} alt="Delete" />
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RouteTable;
