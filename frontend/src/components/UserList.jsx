import React, { useState, useEffect } from "react";
import { deleteAccount } from "./../request/user";

// Données fictives des utilisateurs
const fakeUsersData = [
  { id: 1, username: "User1", email: "user1@example.com", permission: 1 },
  { id: 2, username: "User2", email: "user2@example.com", permission: 1 },
  { id: 3, username: "User3", email: "user3@example.com", permission: 1 },
  // Ajoute plus d'utilisateurs si nécessaire
];

const UserList = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // Ici, tu récupérerais les données réelles des utilisateurs via une requête API
    // Pour l'instant, nous utilisons des données fictives
    setUsers(fakeUsersData);
  }, []);

  const handleDeleteUser = async (userID) => {
    try {
      await deleteAccount(userID);
      console.log(`User with ID ${userID} deleted`);
      // Mise à jour de l'affichage après la suppression
      setUsers(users.filter((user) => user.id !== userID));
    } catch (error) {
      console.error("Error deleting user:", error);
    }
  };

  return (
    <div className="flex overflow-x-auto py-4">
      {users.map((user) => (
        <div
          key={user.id}
          className="flex-none w-60 h-40 m-2 bg-white shadow-lg rounded-lg p-4 flex flex-col justify-between"
        >
          <div>
            <h2 className="font-bold">{user.username}</h2>
            <p>{user.email}</p>
          </div>
          <button
            className="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-4 rounded"
            onClick={() => handleDeleteUser(user.id)}
          >
            Delete
          </button>
        </div>
      ))}
    </div>
  );
};

export default UserList;
