package flags

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flag struct {
	Long string
	Short string
	Usage string
	Value any
}

type Flags map[string]*Flag;

func (flags Flags) Var(long, short string, variable any, msg string) {
	flag := Flag{long, short, msg, variable};
	flags[long] = &flag;
	flags[short] = &flag;
}

func (flags Flags) VarBool(long, short, msg string) *bool {
	s := false;
	flags.Var(long, short, &s, msg);
	return &s;
}

func (flags Flags) Usage() {
	fmt.Printf("Usage: %v\n", os.Args[0]);
	showed := make(map[*Flag]string);
	for _,f := range flags {
		_, ok := showed[f];
		if ok { continue; }
		fmt.Printf("  --%v , -%v\n  \t%v\n", f.Long, f.Short, f.Usage);
		showed[f] = f.Long;
	}
}

func (flags Flags) Parse(argv []string) error {
	for i := 0; i < len(argv); i++ {
		arg := argv[i];

		var err error;

		switch {
		case strings.HasPrefix(arg, "--"):
			s := arg[2:];
			flag := flags[s];
			if flag == nil {
				return fmt.Errorf("unknown flag: %v", arg);
			}

			err = setFlag(flag, s, &i, argv);
			if err != nil {
				return err;
			}
		case strings.HasPrefix(arg, "-"):
			shorts := arg[1:];
			loop: for j := range shorts {
				s := string(shorts[j]);
				flag := flags[s];
				if flag == nil {
					return fmt.Errorf("unknown flag: %v", arg);
				}

				err = setFlag(flag, s, &i, argv);
				if err != nil {
					return err;
				}
				if _, ok := flag.Value.(*bool); !ok {
					break loop;
				}
			}
		}

	}

	return nil;
}

func setFlag(flag *Flag, s string, i* int, argv []string) error {
	switch v := flag.Value.(type) {
	case *string:
		if *i+1 >= len(argv) {
			return fmt.Errorf("flag %v needs an argument", s);
		}

		*i++
		*v = argv[*i];
	case *int:
		if *i+1 >= len(argv) {
			return fmt.Errorf("flag %v needs an argument", s);
		}
		*i++
		conv, err := strconv.Atoi(argv[*i]);
		if err != nil {
			return fmt.Errorf("argument for %v flag needs to be an integer", s);
		}
		*v = conv;
	case *bool:
		*v = true;
	}
	return nil;
}

