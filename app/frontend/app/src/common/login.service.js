const ID_TOKEN_KEY = 'auth_token'
const USER_DATA = 'user_data'

export const getToken = () => {
  return window.localStorage.getItem(ID_TOKEN_KEY)
}

export const getUser = () => {
  let jsonData = window.localStorage.getItem(USER_DATA)
  if (jsonData) {
    try {
      return JSON.parse(jsonData)
    } catch (exc) {
      return {}
    }
  }
  return {}
}

export const saveToken = token => {
  window.localStorage.setItem(ID_TOKEN_KEY, token)
}

export const saveUser = user => {
  window.localStorage.setItem(USER_DATA, JSON.stringify(user))
}

export const destroy = () => {
  window.localStorage.removeItem(ID_TOKEN_KEY)
  window.localStorage.removeItem(USER_DATA)
}

export default { getToken, getUser, saveToken, saveUser, destroy }
