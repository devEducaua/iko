package commands

func Sync() error {
	err := Remove();
	if err != nil {
		return err;
	}

	err = Install();
	if err != nil {
		return err;
	}

	return Update();
}

func Remove() error {
	return nil;
}

func Update() error {
	return nil;
}

func Install() error {

	pkgs := GetPackageList();

	return nil;
}

