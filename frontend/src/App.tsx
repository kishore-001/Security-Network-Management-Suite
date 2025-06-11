import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./pages/login/Login";
import Config from "./pages/config/config";
import Health from "./pages/health/health";
import Log from "./pages/log/log";
import Backup from "./pages/backup/backup";
import Alert from "./pages/alert/alert";
import Resource from "./pages/resource/resource";
import Settings from "./pages/settings/settings";

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Config />} />
        <Route path="/login" element={<Login />} />
        <Route path="/health" element={<Health />} />
        <Route path="/log" element={<Log />} />
        <Route path="/backup" element={<Backup />} />
        <Route path="/alert" element={<Alert />} />
        <Route path="/resource" element={<Resource />} />
        <Route path="/settings" element={<Settings />} />
      </Routes>
    </Router>
  );
}
