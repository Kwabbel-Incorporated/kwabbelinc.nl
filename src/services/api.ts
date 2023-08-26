import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

export const fetchTextContent = async (contentId: number) => {
  try {
    const response = await api.get(`/textcontent/${contentId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};
