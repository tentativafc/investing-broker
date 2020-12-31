import axios from 'axios'
import Vue from 'vue'
import VueAxios from 'vue-axios'
import LoginService from '@/common/login.service'
import { API_URL } from '@/common/config'

const ApiService = {
  init() {
    Vue.use(VueAxios, axios)
    Vue.axios.defaults.baseURL = API_URL
  },

  setHeader() {
    console.log(LoginService.getToken())
    Vue.axios.defaults.headers.common[
      'Authorization'
    ] = `Bearer ${LoginService.getToken()}`
  },

  query(resource, params) {
    return Vue.axios.get(resource, params)
  },

  get(resource, slug = '') {
    return Vue.axios.get(`${resource}/${slug}`)
  },

  post(resource, params) {
    return Vue.axios.post(`${resource}`, params).catch(error => {
      console.log('Erro animal')
      throw new Error(`[RWV] ApiService ${error}`)
    })
  },

  update(resource, slug, params) {
    return Vue.axios.put(`${resource}/${slug}`, params)
  },

  put(resource, params) {
    return Vue.axios.put(`${resource}`, params)
  },

  delete(resource) {
    return Vue.axios.delete(resource)
  }
}

export default ApiService
