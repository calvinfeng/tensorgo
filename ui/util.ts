/**
 * @author Calvin Feng
 */

 import Axios from "axios";


 export async function classifyImageFile(file: File): Promise<any> {
    const formData: FormData = new FormData();
    formData.append("image", file);

    try {
        const response = await Axios.post("api/tf/recognition/", formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });
        return response.data.results;
    } catch(e) {
        throw e.response.data;
    }
 }