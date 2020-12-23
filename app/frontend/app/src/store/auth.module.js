import ApiService from "@/common/api.service";
import LoginService from "@/common/login.service";
import {
  LOGIN,
  LOGOUT,
  REGISTER,
  CHECK_AUTH,
  UPDATE_USER
} from "./actions.type";
import { SET_AUTH, PURGE_AUTH, SET_ERROR } from "./mutations.type";

const initialState = {
  errors: [],
  user: {... LoginService.getUser()},
  isAuthenticated: !!LoginService.getToken()
}

const state = { ... initialState };

const getters = {
  user: (state) =>  state.user,
  isAuthenticated: (state) => state.isAuthenticated
};

const actions = {
  [LOGIN]({commit}, credentials) {
    return new Promise(resolve => {
      ApiService.post("users/login", { ... credentials })
        .then(({ data }) => {
          commit(SET_AUTH, data.user);
          resolve(data);
        })
        .catch(({ response }) => {
          commit(SET_ERROR, response.data.errors);
        });
    });
  },
  [LOGOUT]({commit}) {
    commit(PURGE_AUTH);
  },
  [REGISTER]({commit}, credentials) {
    return new Promise((resolve, reject) => {
      ApiService.post("users", { user: credentials })
        .then(({ data }) => {
          commit(SET_AUTH, data.user);
          resolve(data);
        })
        .catch(({ response }) => {
          commit(SET_ERROR, response.data.errors);
          reject(response);
        });
    });
  },
  [CHECK_AUTH]({commit, state}) {
    if (state.isAuthenticated) {
      ApiService.setHeader();
      ApiService.get("users", state.user.id)
        .then(({ data }) => {
          commit(SET_AUTH, data);
        })
        .catch(({ response }) => {
          commit(SET_ERROR, response.data.errors);
        });
    } else {
      commit(PURGE_AUTH);
    }
  },
  [UPDATE_USER](context, payload) {
    const { email, username, password, image, bio } = payload;
    const user = {
      email,
      username,
      bio,
      image
    };
    if (password) {
      user.password = password;
    }

    return ApiService.put("users", user).then(({ data }) => {
      context.commit(SET_AUTH, data.user);
      return data;
    });
  }
};

const mutations = {
  [SET_ERROR](state, error) {
    state.errors = error;
  },
  [SET_AUTH](state, user) {
    state.isAuthenticated = true;
    state.user = {... user};
    state.errors = {};
    LoginService.saveToken(state.user.auth_token);
    LoginService.saveUser(state.user);
  },
  [PURGE_AUTH](state) {
    state.isAuthenticated = false;
    state.user = {};
    state.errors = {};
    LoginService.destroy();
  }
};

export default {
  state,
  actions,
  mutations,
  getters
};