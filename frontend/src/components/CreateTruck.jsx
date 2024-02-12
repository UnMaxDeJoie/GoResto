import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createTruck } from "./../request/truck";

const CreateTruck = () => {
  const [name, setName] = useState("");
  const [slotBuffer, setSlotBuffer] = useState("");
  const [opening, setOpening] = useState("");
  const [closing, setClosing] = useState("");
  const [userID, setUserID] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const truckData = {
      name,
      slotBuffer: parseInt(slotBuffer, 10),
      opening,
      closing,
      userID: parseInt(userID, 10),
    };

    try {
      await createTruck(truckData);
      console.log("Truck created successfully");
    } catch (error) {
      console.error("There was an error creating the truck:", error);
    }
  };

  return (
    <form
      className=" flex flex-col justify-center items-center"
      onSubmit={handleSubmit}
    >
      <p className="mb-5 text-3xl uppercase text-gray-600">Create Truck</p>
      <input
        type="text"
        name="name"
        id="name"
        placeholder="Truck Name"
        className="mb-4 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
        autoComplete="off"
        required
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <input
        type="number"
        name="slotBuffer"
        id="slotBuffer"
        placeholder="Slot Buffer"
        className="mb-4 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
        required
        value={slotBuffer}
        onChange={(e) => setSlotBuffer(e.target.value)}
      />
      <input
        type="time"
        name="opening"
        id="opening"
        placeholder="Opening Time"
        className="mb-4 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
        required
        value={opening}
        onChange={(e) => setOpening(e.target.value)}
      />
      <input
        type="time"
        name="closing"
        id="closing"
        placeholder="Closing Time"
        className="mb-4 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
        required
        value={closing}
        onChange={(e) => setClosing(e.target.value)}
      />
      <input
        type="number"
        name="userID"
        id="userID"
        placeholder="User ID"
        className="mb-6 p-3 w-80 focus:border-purple-700 rounded border-2 outline-none"
        required
        value={userID}
        onChange={(e) => setUserID(e.target.value)}
      />
      <button
        type="submit"
        className="bg-purple-600 hover:bg-purple-900 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-80"
      >
        Create Truck
      </button>
    </form>
  );
};

export default CreateTruck;
