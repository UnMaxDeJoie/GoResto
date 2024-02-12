import { NavLink, useNavigate } from "react-router-dom";
import { useState } from "react";

const Register = () => {
  const [userName, setUserName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [userRole, setUserRole] = useState(0);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    // if (email && password && confirmPassword && userName && userRole) {
    //     if (password === confirmPassword) {
    //         try {
    //             // Assurez-vous que les noms des clés correspondent à ceux attendus par votre API
    //             const response = await Signup(userName, email, password, userRole);
    //             console.log('Registration successful', response);
    //             localStorage.setItem('id', response.id);
    //             navigate('/login');
    //         } catch (error) {
    //             console.error('Registration failed', error);
    //         }
    //     } else {
    //         console.log('Passwords do not match');
    //     }
    // } else {
    //     console.log('Please fill in all fields correctly');
    // }

    // fake login for now

    navigate("/login");
  };

  return (
    <div className="w-full flex justify-center items-center h-screen bg-gray-100">
      <form
        className="p-10 bg-white rounded flex justify-center items-center flex-col shadow-md"
        onSubmit={handleSubmit}
      >
        <p className="mb-5 text-3xl uppercase text-gray-600">Register</p>
        <input
          type="text"
          name="userName"
          id="userName"
          placeholder="User Name"
          className="mb-5 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
          autoComplete="off"
          required
          value={userName}
          onChange={(e) => setUserName(e.target.value)}
        />
        <input
          type="email"
          name="email"
          id="email"
          placeholder="Email"
          className="mb-5 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
          autoComplete="off"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          name="password"
          id="password"
          placeholder="Password"
          className="mb-5 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
          required
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <input
          type="password"
          name="confirmPassword"
          id="confirmPassword"
          placeholder="Confirm Password"
          className="mb-5 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
          required
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
        <input
          type="text"
          name="userRole"
          id="userRole"
          placeholder="User Role"
          className="mb-5 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
          autoComplete="off"
          required
          value={userRole}
          onChange={(e) => setUserRole(e.target.value)}
        />
        <button
          type="submit"
          className="bg-purple-600 hover:bg-purple-900 text-white font-bold p-2 rounded w-80"
        >
          Register
        </button>
        <div className="mt-5">
          Already have an account?{" "}
          <NavLink to="/login" className="text-blue-600 hover:underline">
            Log in
          </NavLink>
        </div>
      </form>
    </div>
  );
};

export default Register;
