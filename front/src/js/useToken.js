import React, {useState} from 'react'
import constants from '../constants.js'

function useToken() {
    const getToken = () => {
        const tokenString = localStorage.getItem(constants.AUTH_TOKEN)
        const userToken = JSON.parse(tokenString);
        return userToken
    }

    const [token, setToken] = useState(getToken())

    const saveToken = (userToken) => {
        localStorage.setItem(constants.AUTH_TOKEN, JSON.stringify(userToken))
        setToken(userToken)
    }

    return {
        setToken: saveToken,
        token
    }
}

export default useToken