import { API_URL } from "./api";

export const createTruck = async (truckData) => {
    try {
        const response = await fetch(`${API_URL}/truck`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(truckData),
        });
        if (!response.ok) {
            throw new Error("Error creating truck");
        }
        return await response.json();
    } catch (error) {
        console.error("There was an error creating the truck:", error);
        throw error;
    }
};

export const getTruckById = async (id) => {
    const response = await fetch(`${API_URL}/truck/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting truck");
    }
    return response.json();
}


export const deleteTruck = async (id) => {
    const response = await fetch(`${API_URL}/truck/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting truck");
    }
    return response.json();
}

export const GetAvailableTrucks = async () => {
    const response = await fetch(`${API_URL}/trucks`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting trucks");
    }
    return response.json();
}