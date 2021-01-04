import ApiService from '@/common/api.service'
import StorageService from '@/common/storage.service'
import router from '../router'

import {
  LOGIN,
  LOGOUT,
  REGISTER,
  CHECK_AUTH,
  UPDATE_USER,
  FORGET_PASSWORD
} from './actions.type'
import { SET_AUTH, PURGE_AUTH, SET_ERROR } from './mutations.type'

const initialState = {
  errors: [],
  user: { ...StorageService.getUser() },
  isAuthenticated: !!StorageService.getToken()
}

const state = { ...initialState }

const getters = {
  user: state => state.user,
  isAuthenticated: state => state.isAuthenticated
}

const actions = {
  [LOGIN]({ commit }, credentials) {
    return new Promise((resolve, reject) => {
      ApiService.post('users/login', { ...credentials })
        .then(({ data }) => {
          commit(SET_AUTH, data)
          resolve(data)
        })
        .catch(({ response }) => {
          let error = ''
          if (response && response.data) {
            error = response.data
          }

          commit(SET_ERROR, error)
          reject(error)
        })
    })
  },
  [LOGOUT]({ commit }) {
    commit(PURGE_AUTH)
  },
  [REGISTER]({ commit }, payload) {
    return new Promise((resolve, reject) => {
      ApiService.post('users', { ...payload })
        .then(({ data }) => {
          commit(SET_AUTH, data)
          resolve(data)
        })
        .catch(({ response }) => {
          let error = null
          if (response && response.data) {
            error = response.data
          }
          commit(SET_ERROR, error)
          reject(error)
        })
    })
  },
  [FORGET_PASSWORD]({ commit }, payload) {
    return new Promise((resolve, reject) => {
      ApiService.post('users/recover', { ...payload })
        .then(({ data }) => {
          resolve(data)
        })
        .catch(({ response }) => {
          let error = null
          if (response && response.data) {
            error = response.data
          }
          commit(SET_ERROR, error)
          reject(error)
        })
    })
  },
  [CHECK_AUTH]({ commit, state }) {
    if (state.isAuthenticated) {
      ApiService.setHeader()
      ApiService.get('users', state.user.id)
        .then(({ data }) => {
          commit(SET_AUTH, data)
        })
        .catch(({ response }) => {
          let error = null
          if (response && response.data) {
            error = response.data
          }
          commit(SET_ERROR, error)
        })
    } else {
      commit(PURGE_AUTH)
    }
  },
  [UPDATE_USER](context, payload) {
    const { email, username, password, image, bio } = payload
    const user = {
      email,
      username,
      bio,
      image
    }
    if (password) {
      user.password = password
    }

    return ApiService.put('users', user).then(({ data }) => {
      context.commit(SET_AUTH, data.user)
      return data
    })
  }
}

const mutations = {
  [SET_AUTH](state, loginData) {
    state.isAuthenticated = true
    state.user = { ...loginData }
    state.errors = []
    StorageService.saveToken(loginData.auth_token)
    StorageService.saveUser(state.user)
  },
  [PURGE_AUTH](state) {
    state.isAuthenticated = false
    state.user = {}
    state.errors = []
    StorageService.destroy()
    router.go('/')
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}
