import React, { useState } from "react";
import { Link } from "react-router-dom";

export const Home = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const handleLogout = () => {
    setIsLoggedIn(false);
  };

  return (
    <div className="bg-gray-100 min-h-screen">
      <nav className="bg-blue-500 p-4">
        <div className="container mx-auto flex justify-between items-center">
          <Link to="/" className="text-white text-2xl font-bold">
            Habbats Institute
          </Link>
          <div>
            {isLoggedIn ? (
              <button onClick={handleLogout} className="text-white mx-4">
                Logout
              </button>
            ) : (
              <Link to="/register" className="text-white mx-4">
                Register
              </Link>
            )}
          </div>
        </div>
      </nav>

      <div className="flex justify-center items-center h-screen">
        <h1 className="text-5xl text-red-500">Ini adalah home</h1>
      </div>
    </div>
  );
};
