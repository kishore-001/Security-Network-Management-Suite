import "./App.css";
import { Routes, Route } from "react-router-dom";
import Login from "./pages/login/login";
export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
    </Routes>
  );
}
