import Axios from "axios";

let store, next = null

export default () => {
  return next || new Promise(first_resolve => {
    return Axios.get('/api/sites')
      .then(res => res.data)
      .then(sites => {
        first_resolve(store = {
          state: {
            data: sites
          }
        })
        return store
      })
  })
}