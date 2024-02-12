import React, { useState, useEffect } from "react";

const EditTruckForm = ({ selectedTruckData, onSave }) => {
  const [name, setName] = useState("");
  const [slotBuffer, setSlotBuffer] = useState("");
  const [opening, setOpening] = useState("");
  const [closing, setClosing] = useState("");
  const [userID, setUserID] = useState("");

  useEffect(() => {
    if (selectedTruckData) {
      setName(selectedTruckData.name);
      setSlotBuffer(selectedTruckData.slotBuffer);
      setOpening(selectedTruckData.opening);
      setClosing(selectedTruckData.closing);
      setUserID(selectedTruckData.userID);
    }
  }, [selectedTruckData]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    onSave({
      id: selectedTruckData.id,
      name,
      slotBuffer,
      opening,
      closing,
      userID,
    });
  };

  return selectedTruckData ? (
    <form className="p-4 bg-white rounded shadow-md" onSubmit={handleSubmit}>
      <h2 className="text-xl font-bold mb-4">Edit Truck</h2>
      <div className="flex flex-col mb-4">
        <label htmlFor="name" className="mb-2">
          Truck Name:
        </label>
        <input
          type="text"
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="border-2 rounded p-2"
        />
      </div>
      <button
        type="submit"
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        Save Changes
      </button>
    </form>
  ) : null;
};

export default EditTruckForm;
