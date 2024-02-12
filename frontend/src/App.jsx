import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Navbar from "./components/layout/Navbar";
import Login from "./pages/Login";
import Register from "./pages/Register";
import AdminPage from "./pages/AdminPage";
import ManagerPage from "./pages/ManagerPage";
import CustomerPage from "./pages/CustomerPage";

import "./App.css";
function App() {
  return (
    <Router>
      <div className="w-full h-full flex flex-col justify-start items-center">
        <Navbar />
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/customer_page" element={<CustomerPage />} />
          <Route path="/manager_page" element={<ManagerPage />} />
          <Route path="/admin_page" element={<AdminPage />} />
          <Route
            path="*"
            element={
              <div className="h-full w-full flex flex-col items-center justify-center gap-4">
                <h1>Not Found</h1>
              </div>
            }
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
