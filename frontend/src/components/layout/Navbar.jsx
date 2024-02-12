import { NavLink } from "react-router-dom";

const Navbar = () => {
  const handleLogout = () => {
    localStorage.clear();
  };

  return (
    <nav className="bg-gray-800 w-full p-4 text-white flex flex-row items-center justify-between">
      <div className="flex justify-center space-x-4 items-center">
        <h1 className="text-2xl font-bold">Go Resto</h1>
        <NavLink
          to="/login"
          activeClassName="active"
          className="text-white hover:text-gray-300"
        >
          Login
        </NavLink>
        <NavLink
          to="/register"
          activeClassName="active"
          className="text-white hover:text-gray-300"
        >
          Register
        </NavLink>
      </div>
      <div className="flex justify-center space-x-4 items-center">
        <NavLink
          to="/customer_page"
          activeClassName="active"
          className="text-white hover:text-gray-300"
        >
          Customer
        </NavLink>
        <NavLink
          to="/manager_page"
          activeClassName="active"
          className="text-white hover:text-gray-300"
        >
          Manager
        </NavLink>
        <NavLink
          to="/admin_page"
          activeClassName="active"
          className="text-white hover:text-gray-300"
        >
          Admin
        </NavLink>
        <button
          className="bg-red-500 px-4 py-2 rounded-md hover:bg-red-600"
          onClick={handleLogout}
        >
          logout
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
