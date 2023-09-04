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

export const fetchLatestArticles = async (latests: number) => {
  try {
    const response = await api.get(`/articles/latests/${latests}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const fetchMedia = async (mediaId: number) => {
  try {
    const response = await api.get(`/media/${mediaId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};
