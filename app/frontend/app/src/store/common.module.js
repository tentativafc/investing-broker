import { SET_GENERIC_ERROR, SET_ERROR } from './mutations.type'

const initialState = {
  generic_errors: [],
  errors: []
}

const state = { ...initialState }

const getters = {}

const actions = {}

const mutations = {
  [SET_GENERIC_ERROR](state, error) {
    if (!error) {
      error = { code: 500, message: 'Sistema indisponível' }
    }
    state.generic_errors = [...state.generic_errors, error]
  },
  [SET_ERROR](state, error_message) {
    if (!error_message) {
      error_message = 'Sistema indisponível'
    }
    state.errors = [...state.errors, error_message]
  }
}

export default {
  state,
  actions,
  mutations,
  getters
}
