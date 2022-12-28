package serialmedia

//func (wrapper *SerialMediaWrapper) GetVersion(ctx context.Context, target string) (serialMediaControl.EndpointVersion, error) {
//	wrapper.OpLock.RLock()
//	defer wrapper.OpLock.RUnlock()
//
//	connection, hasConnection := wrapper.SerialMediaCons[target]
//	if !hasConnection {
//		return serialMediaControl.EndpointVersion{}, errors.New("endpoint not found")
//	}
//
//	return connection.GetVersion(ctx)
//}
