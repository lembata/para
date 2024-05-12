//import fetch from 'node-fetch';
//import { promises as fs } from 'fs';

//TODO: Load from env file
const url = 'http://localhost:8080/';

const API = {
  Accounts: {
    add: async (data) => {
      console.log("Accounts.add", data);
      return await requester._post('api/accounts/add', data)
    },
    edit: async (id, data) => {
      console.log("Accounts.edit", data);
      return await requester._post(`api/accounts/edit/${id}`, data)
    },
    get: async (id) => {
      console.log("Accounts.get", id);
      return await requester._get(`api/accounts/${id}`)
    },
  }
}

const requester = {

  _get: async (slug) => {
    return await fetch(`${url}${slug}`)
      .then(response => response.json())
      .catch(error => ({ success: false, error: error, data: null }));
  },

  _post: async (slug, data) => {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    }

    return await fetch(`${url}${slug}`, requestOptions)
      .then(response => response.json())
      .catch(error => ({ success: false, error: error, data: null }));
  }
}

export default API;