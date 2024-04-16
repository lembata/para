package api

import (
	"net/http"
)

type SessionCache struct {
}

func (s *SessionCache) GetSession(w *http.ResponseWriter,r *http.Request) {
	//TOGO: check if session exists
} 

 func authenticateHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//next.ServeHTTP();
			w.Header().Add("WWW-Authenticate", "FormBased")
			w.WriteHeader(http.StatusUnauthorized)
			return
			// c := config.GetInstance()
			//
			// // error if external access tripwire activated
			// if accessErr := session.CheckExternalAccessTripwire(c); accessErr != nil {
			// 	http.Error(w, tripwireActivatedErrMsg, http.StatusForbidden)
			// 	return
			// }
			//
			// userID, err := manager.GetInstance().SessionStore.Authenticate(w, r)
			// if err != nil {
			// 	if errors.Is(err, session.ErrUnauthorized) {
			// 		http.Error(w, err.Error(), http.StatusInternalServerError)
			// 		return
			// 	}
			//
			// 	// unauthorized error
			// 	w.Header().Add("WWW-Authenticate", "FormBased")
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }
			//
			// if err := session.CheckAllowPublicWithoutAuth(c, r); err != nil {
			// 	var accessErr session.ExternalAccessError
			// 	if errors.As(err, &accessErr) {
			// 		session.LogExternalAccessError(accessErr)
			//
			// 		err := c.ActivatePublicAccessTripwire(net.IP(accessErr).String())
			// 		if err != nil {
			// 			logger.Errorf("Error activating public access tripwire: %v", err)
			// 		}
			//
			// 		http.Error(w, externalAccessErrMsg, http.StatusForbidden)
			// 	} else {
			// 		logger.Errorf("Error checking external access security: %v", err)
			// 		w.WriteHeader(http.StatusInternalServerError)
			// 	}
			// 	return
			// }
			//
			// ctx := r.Context()
			//
			// if c.HasCredentials() {
			// 	// authentication is required
			// 	if userID == "" && !allowUnauthenticated(r) {
			// 		// if graphql or a non-webpage was requested, we just return a forbidden error
			// 		ext := path.Ext(r.URL.Path)
			// 		if r.URL.Path == gqlEndpoint || (ext != "" && ext != ".html") {
			// 			w.Header().Add("WWW-Authenticate", "FormBased")
			// 			w.WriteHeader(http.StatusUnauthorized)
			// 			return
			// 		}
			//
			// 		prefix := getProxyPrefix(r)
			//
			// 		// otherwise redirect to the login page
			// 		returnURL := url.URL{
			// 			Path:     prefix + r.URL.Path,
			// 			RawQuery: r.URL.RawQuery,
			// 		}
			// 		q := make(url.Values)
			// 		q.Set(returnURLParam, returnURL.String())
			// 		u := url.URL{
			// 			Path:     prefix + loginEndpoint,
			// 			RawQuery: q.Encode(),
			// 		}
			// 		http.Redirect(w, r, u.String(), http.StatusFound)
			// 		return
			// 	}
			// }
			//
			// ctx = session.SetCurrentUserID(ctx, userID)
			//
			// r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
    
