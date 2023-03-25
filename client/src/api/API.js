export const baseURL = 'http://localhost:8080/api';

//FOR FUTURE OPTIMIZATION
// export const defaultHeaders = {
//     Accept: "application/json",
//     Credentials: "include",
// };

// export async function fetchWithCredentials(url, options) {
//     const response = await fetch(url, {
//       ...options,
//       headers: {
//         ...defaultHeaders,
//         ...options.headers,
//       },
//       credentials: "include",
//     });
  
//     if (!response.ok) {
//       throw new Error(
//         `Could not fetch ${url}. Status: ${response.status} ${response.statusText}`
//       );
//     }
  
//     return response.json();
// }