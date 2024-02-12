import { API_URL } from "./api";

export const createChart = async (chartData) => {
    try {
        const response = await fetch(`${API_URL}/chart`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(chartData),
        });
        if (!response.ok) {
            throw new Error("Error creating chart");
        }
        return await response.json();
    } catch (error) {
        console.error("There was an error creating the chart:", error);
        throw error;
    }
};
export const getChartById = async (id) => {
    const response = await fetch(`${API_URL}/chart/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting chart");
    }
    return response.json();
}
export const getChartsByTruck = async (idTruck) => {
    const response = await fetch(`${API_URL}/chart/truck/${idTruck}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting charts");
    }
    return response.json();
}
export const updateChart = async (id) => {
    const response = await fetch(`${API_URL}/chart/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting chart");
    }
    return response.json();
};
export const deleteChart = async (consumableId) => {
    const response = await fetch(`${API_URL}/chart/consumable/${consumableId}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Error getting chart");
    }
    return response.json();
};