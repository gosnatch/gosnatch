package gosnatch

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assets_js_ds_store reads file data from disk. It returns an error on failure.
func assets_js_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/.DS_Store"
	name := "assets/js/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_bootstrap_js reads file data from disk. It returns an error on failure.
func assets_js_bootstrap_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/bootstrap.js"
	name := "assets/js/bootstrap.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_js_bootstrap_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/bootstrap.min.js"
	name := "assets/js/bootstrap.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_de_js reads file data from disk. It returns an error on failure.
func assets_js_de_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/de.js"
	name := "assets/js/de.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_jquery_knob_min_js reads file data from disk. It returns an error on failure.
func assets_js_jquery_knob_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/jquery.knob.min.js"
	name := "assets/js/jquery.knob.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_main_js reads file data from disk. It returns an error on failure.
func assets_js_main_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/main.js"
	name := "assets/js/main.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_ds_store reads file data from disk. It returns an error on failure.
func assets_js_external_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/.DS_Store"
	name := "assets/js/external/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_bootstrap3_typeahead_min_js reads file data from disk. It returns an error on failure.
func assets_js_external_bootstrap3_typeahead_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/bootstrap3-typeahead.min.js"
	name := "assets/js/external/bootstrap3-typeahead.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_circle_progress_js reads file data from disk. It returns an error on failure.
func assets_js_external_circle_progress_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/circle-progress.js"
	name := "assets/js/external/circle-progress.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_datatables_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_js_external_datatables_bootstrap_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/dataTables.bootstrap.min.js"
	name := "assets/js/external/dataTables.bootstrap.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_fullcalendar_min_js reads file data from disk. It returns an error on failure.
func assets_js_external_fullcalendar_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/fullcalendar.min.js"
	name := "assets/js/external/fullcalendar.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_jquery_datatables_min_js reads file data from disk. It returns an error on failure.
func assets_js_external_jquery_datatables_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/jquery.dataTables.min.js"
	name := "assets/js/external/jquery.dataTables.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_jquery_dynatable_js reads file data from disk. It returns an error on failure.
func assets_js_external_jquery_dynatable_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/jquery.dynatable.js"
	name := "assets/js/external/jquery.dynatable.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_moment_min_js reads file data from disk. It returns an error on failure.
func assets_js_external_moment_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/moment.min.js"
	name := "assets/js/external/moment.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_external_ohsnap_js reads file data from disk. It returns an error on failure.
func assets_js_external_ohsnap_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/external/ohsnap.js"
	name := "assets/js/external/ohsnap.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_vendor_ds_store reads file data from disk. It returns an error on failure.
func assets_js_vendor_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/vendor/.DS_Store"
	name := "assets/js/vendor/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_vendor_jquery_ba_throttle_debounce_min_js reads file data from disk. It returns an error on failure.
func assets_js_vendor_jquery_ba_throttle_debounce_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/vendor/jquery.ba-throttle-debounce.min.js"
	name := "assets/js/vendor/jquery.ba-throttle-debounce.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_vendor_jquery_min_js reads file data from disk. It returns an error on failure.
func assets_js_vendor_jquery_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/vendor/jquery.min.js"
	name := "assets/js/vendor/jquery.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_vendor_underscore_min_js reads file data from disk. It returns an error on failure.
func assets_js_vendor_underscore_min_js() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/js/vendor/underscore-min.js"
	name := "assets/js/vendor/underscore-min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_ds_store reads file data from disk. It returns an error on failure.
func assets_css_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/.DS_Store"
	name := "assets/css/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_css_bootstrap_min_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/bootstrap.min.css"
	name := "assets/css/bootstrap.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_bootswatch_css reads file data from disk. It returns an error on failure.
func assets_css_bootswatch_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/bootswatch.css"
	name := "assets/css/bootswatch.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_bootswatch_less reads file data from disk. It returns an error on failure.
func assets_css_bootswatch_less() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/bootswatch.less"
	name := "assets/css/bootswatch.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_datatables_bootstrap_css reads file data from disk. It returns an error on failure.
func assets_css_datatables_bootstrap_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/dataTables.bootstrap.css"
	name := "assets/css/dataTables.bootstrap.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_font_awesome_min_css reads file data from disk. It returns an error on failure.
func assets_css_font_awesome_min_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/font-awesome.min.css"
	name := "assets/css/font-awesome.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fullcalendar_min_css reads file data from disk. It returns an error on failure.
func assets_css_fullcalendar_min_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/fullcalendar.min.css"
	name := "assets/css/fullcalendar.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_jquery_dynatable_css reads file data from disk. It returns an error on failure.
func assets_css_jquery_dynatable_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/jquery.dynatable.css"
	name := "assets/css/jquery.dynatable.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_mixins_less reads file data from disk. It returns an error on failure.
func assets_css_mixins_less() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/mixins.less"
	name := "assets/css/mixins.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_style_css reads file data from disk. It returns an error on failure.
func assets_css_style_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/style.css"
	name := "assets/css/style.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_style_less reads file data from disk. It returns an error on failure.
func assets_css_style_less() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/style.less"
	name := "assets/css/style.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_variables_css reads file data from disk. It returns an error on failure.
func assets_css_variables_css() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/variables.css"
	name := "assets/css/variables.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_variables_less reads file data from disk. It returns an error on failure.
func assets_css_variables_less() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/css/variables.less"
	name := "assets/css/variables.less"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_79349_jpg reads file data from disk. It returns an error on failure.
func assets_images_79349_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/79349.jpg"
	name := "assets/images/79349.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_ahs_jpg reads file data from disk. It returns an error on failure.
func assets_images_ahs_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/ahs.jpg"
	name := "assets/images/ahs.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_banshee_jpg reads file data from disk. It returns an error on failure.
func assets_images_banshee_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/banshee.jpg"
	name := "assets/images/banshee.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_chuck_jpg reads file data from disk. It returns an error on failure.
func assets_images_chuck_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/chuck.jpg"
	name := "assets/images/chuck.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_dexter_jpg reads file data from disk. It returns an error on failure.
func assets_images_dexter_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/dexter.jpg"
	name := "assets/images/dexter.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_elementary_jpg reads file data from disk. It returns an error on failure.
func assets_images_elementary_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/elementary.jpg"
	name := "assets/images/elementary.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_images_suits_jpg reads file data from disk. It returns an error on failure.
func assets_images_suits_jpg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/images/suits.jpg"
	name := "assets/images/suits.jpg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_ds_store reads file data from disk. It returns an error on failure.
func assets_templates_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/.DS_Store"
	name := "assets/templates/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_header_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_header_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/_header.tmpl"
	name := "assets/templates/_header.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_script_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_script_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/_script.tmpl"
	name := "assets/templates/_script.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_addseries_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_addseries_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/addSeries.tmpl"
	name := "assets/templates/addSeries.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_calendar_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_calendar_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/calendar.tmpl"
	name := "assets/templates/calendar.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_history_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_history_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/history.tmpl"
	name := "assets/templates/history.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_index_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_index_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/index.tmpl"
	name := "assets/templates/index.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_presets_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_presets_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/presets.tmpl"
	name := "assets/templates/presets.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_settings_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_settings_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/settings.tmpl"
	name := "assets/templates/settings.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_show_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_show_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/show.tmpl"
	name := "assets/templates/show.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_templates_shows_tmpl reads file data from disk. It returns an error on failure.
func assets_templates_shows_tmpl() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/templates/shows.tmpl"
	name := "assets/templates/shows.tmpl"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_otf reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_otf() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/FontAwesome.otf"
	name := "assets/fonts/FontAwesome.otf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_webfont_eot reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_webfont_eot() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/fontawesome-webfont.eot"
	name := "assets/fonts/fontawesome-webfont.eot"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_webfont_svg reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_webfont_svg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/fontawesome-webfont.svg"
	name := "assets/fonts/fontawesome-webfont.svg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_webfont_ttf reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_webfont_ttf() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/fontawesome-webfont.ttf"
	name := "assets/fonts/fontawesome-webfont.ttf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_webfont_woff reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_webfont_woff() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/fontawesome-webfont.woff"
	name := "assets/fonts/fontawesome-webfont.woff"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_fontawesome_webfont_woff2 reads file data from disk. It returns an error on failure.
func assets_fonts_fontawesome_webfont_woff2() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/fontawesome-webfont.woff2"
	name := "assets/fonts/fontawesome-webfont.woff2"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_eot() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/glyphicons-halflings-regular.eot"
	name := "assets/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_svg() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/glyphicons-halflings-regular.svg"
	name := "assets/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_ttf() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/glyphicons-halflings-regular.ttf"
	name := "assets/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_woff() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/glyphicons-halflings-regular.woff"
	name := "assets/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_fonts_glyphicons_halflings_regular_woff2() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/fonts/glyphicons-halflings-regular.woff2"
	name := "assets/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_translations_ds_store reads file data from disk. It returns an error on failure.
func assets_translations_ds_store() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/translations/.DS_Store"
	name := "assets/translations/.DS_Store"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_translations_de_de_all_json reads file data from disk. It returns an error on failure.
func assets_translations_de_de_all_json() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/translations/de-DE.all.json"
	name := "assets/translations/de-DE.all.json"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_translations_en_us_all_json reads file data from disk. It returns an error on failure.
func assets_translations_en_us_all_json() (*asset, error) {
	path := "/Users/workstation/golang/src/github.com/gosnatch/gosnatch/assets/translations/en-US.all.json"
	name := "assets/translations/en-US.all.json"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/js/.DS_Store": assets_js_ds_store,
	"assets/js/bootstrap.js": assets_js_bootstrap_js,
	"assets/js/bootstrap.min.js": assets_js_bootstrap_min_js,
	"assets/js/de.js": assets_js_de_js,
	"assets/js/jquery.knob.min.js": assets_js_jquery_knob_min_js,
	"assets/js/main.js": assets_js_main_js,
	"assets/js/external/.DS_Store": assets_js_external_ds_store,
	"assets/js/external/bootstrap3-typeahead.min.js": assets_js_external_bootstrap3_typeahead_min_js,
	"assets/js/external/circle-progress.js": assets_js_external_circle_progress_js,
	"assets/js/external/dataTables.bootstrap.min.js": assets_js_external_datatables_bootstrap_min_js,
	"assets/js/external/fullcalendar.min.js": assets_js_external_fullcalendar_min_js,
	"assets/js/external/jquery.dataTables.min.js": assets_js_external_jquery_datatables_min_js,
	"assets/js/external/jquery.dynatable.js": assets_js_external_jquery_dynatable_js,
	"assets/js/external/moment.min.js": assets_js_external_moment_min_js,
	"assets/js/external/ohsnap.js": assets_js_external_ohsnap_js,
	"assets/js/vendor/.DS_Store": assets_js_vendor_ds_store,
	"assets/js/vendor/jquery.ba-throttle-debounce.min.js": assets_js_vendor_jquery_ba_throttle_debounce_min_js,
	"assets/js/vendor/jquery.min.js": assets_js_vendor_jquery_min_js,
	"assets/js/vendor/underscore-min.js": assets_js_vendor_underscore_min_js,
	"assets/css/.DS_Store": assets_css_ds_store,
	"assets/css/bootstrap.min.css": assets_css_bootstrap_min_css,
	"assets/css/bootswatch.css": assets_css_bootswatch_css,
	"assets/css/bootswatch.less": assets_css_bootswatch_less,
	"assets/css/dataTables.bootstrap.css": assets_css_datatables_bootstrap_css,
	"assets/css/font-awesome.min.css": assets_css_font_awesome_min_css,
	"assets/css/fullcalendar.min.css": assets_css_fullcalendar_min_css,
	"assets/css/jquery.dynatable.css": assets_css_jquery_dynatable_css,
	"assets/css/mixins.less": assets_css_mixins_less,
	"assets/css/style.css": assets_css_style_css,
	"assets/css/style.less": assets_css_style_less,
	"assets/css/variables.css": assets_css_variables_css,
	"assets/css/variables.less": assets_css_variables_less,
	"assets/images/79349.jpg": assets_images_79349_jpg,
	"assets/images/ahs.jpg": assets_images_ahs_jpg,
	"assets/images/banshee.jpg": assets_images_banshee_jpg,
	"assets/images/chuck.jpg": assets_images_chuck_jpg,
	"assets/images/dexter.jpg": assets_images_dexter_jpg,
	"assets/images/elementary.jpg": assets_images_elementary_jpg,
	"assets/images/suits.jpg": assets_images_suits_jpg,
	"assets/templates/.DS_Store": assets_templates_ds_store,
	"assets/templates/_header.tmpl": assets_templates_header_tmpl,
	"assets/templates/_script.tmpl": assets_templates_script_tmpl,
	"assets/templates/addSeries.tmpl": assets_templates_addseries_tmpl,
	"assets/templates/calendar.tmpl": assets_templates_calendar_tmpl,
	"assets/templates/history.tmpl": assets_templates_history_tmpl,
	"assets/templates/index.tmpl": assets_templates_index_tmpl,
	"assets/templates/presets.tmpl": assets_templates_presets_tmpl,
	"assets/templates/settings.tmpl": assets_templates_settings_tmpl,
	"assets/templates/show.tmpl": assets_templates_show_tmpl,
	"assets/templates/shows.tmpl": assets_templates_shows_tmpl,
	"assets/fonts/FontAwesome.otf": assets_fonts_fontawesome_otf,
	"assets/fonts/fontawesome-webfont.eot": assets_fonts_fontawesome_webfont_eot,
	"assets/fonts/fontawesome-webfont.svg": assets_fonts_fontawesome_webfont_svg,
	"assets/fonts/fontawesome-webfont.ttf": assets_fonts_fontawesome_webfont_ttf,
	"assets/fonts/fontawesome-webfont.woff": assets_fonts_fontawesome_webfont_woff,
	"assets/fonts/fontawesome-webfont.woff2": assets_fonts_fontawesome_webfont_woff2,
	"assets/fonts/glyphicons-halflings-regular.eot": assets_fonts_glyphicons_halflings_regular_eot,
	"assets/fonts/glyphicons-halflings-regular.svg": assets_fonts_glyphicons_halflings_regular_svg,
	"assets/fonts/glyphicons-halflings-regular.ttf": assets_fonts_glyphicons_halflings_regular_ttf,
	"assets/fonts/glyphicons-halflings-regular.woff": assets_fonts_glyphicons_halflings_regular_woff,
	"assets/fonts/glyphicons-halflings-regular.woff2": assets_fonts_glyphicons_halflings_regular_woff2,
	"assets/translations/.DS_Store": assets_translations_ds_store,
	"assets/translations/de-DE.all.json": assets_translations_de_de_all_json,
	"assets/translations/en-US.all.json": assets_translations_en_us_all_json,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			".DS_Store": &_bintree_t{assets_css_ds_store, map[string]*_bintree_t{
			}},
			"bootstrap.min.css": &_bintree_t{assets_css_bootstrap_min_css, map[string]*_bintree_t{
			}},
			"bootswatch.css": &_bintree_t{assets_css_bootswatch_css, map[string]*_bintree_t{
			}},
			"bootswatch.less": &_bintree_t{assets_css_bootswatch_less, map[string]*_bintree_t{
			}},
			"dataTables.bootstrap.css": &_bintree_t{assets_css_datatables_bootstrap_css, map[string]*_bintree_t{
			}},
			"font-awesome.min.css": &_bintree_t{assets_css_font_awesome_min_css, map[string]*_bintree_t{
			}},
			"fullcalendar.min.css": &_bintree_t{assets_css_fullcalendar_min_css, map[string]*_bintree_t{
			}},
			"jquery.dynatable.css": &_bintree_t{assets_css_jquery_dynatable_css, map[string]*_bintree_t{
			}},
			"mixins.less": &_bintree_t{assets_css_mixins_less, map[string]*_bintree_t{
			}},
			"style.css": &_bintree_t{assets_css_style_css, map[string]*_bintree_t{
			}},
			"style.less": &_bintree_t{assets_css_style_less, map[string]*_bintree_t{
			}},
			"variables.css": &_bintree_t{assets_css_variables_css, map[string]*_bintree_t{
			}},
			"variables.less": &_bintree_t{assets_css_variables_less, map[string]*_bintree_t{
			}},
		}},
		"fonts": &_bintree_t{nil, map[string]*_bintree_t{
			"FontAwesome.otf": &_bintree_t{assets_fonts_fontawesome_otf, map[string]*_bintree_t{
			}},
			"fontawesome-webfont.eot": &_bintree_t{assets_fonts_fontawesome_webfont_eot, map[string]*_bintree_t{
			}},
			"fontawesome-webfont.svg": &_bintree_t{assets_fonts_fontawesome_webfont_svg, map[string]*_bintree_t{
			}},
			"fontawesome-webfont.ttf": &_bintree_t{assets_fonts_fontawesome_webfont_ttf, map[string]*_bintree_t{
			}},
			"fontawesome-webfont.woff": &_bintree_t{assets_fonts_fontawesome_webfont_woff, map[string]*_bintree_t{
			}},
			"fontawesome-webfont.woff2": &_bintree_t{assets_fonts_fontawesome_webfont_woff2, map[string]*_bintree_t{
			}},
			"glyphicons-halflings-regular.eot": &_bintree_t{assets_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
			}},
			"glyphicons-halflings-regular.svg": &_bintree_t{assets_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
			}},
			"glyphicons-halflings-regular.ttf": &_bintree_t{assets_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
			}},
			"glyphicons-halflings-regular.woff": &_bintree_t{assets_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
			}},
			"glyphicons-halflings-regular.woff2": &_bintree_t{assets_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
			}},
		}},
		"images": &_bintree_t{nil, map[string]*_bintree_t{
			"79349.jpg": &_bintree_t{assets_images_79349_jpg, map[string]*_bintree_t{
			}},
			"ahs.jpg": &_bintree_t{assets_images_ahs_jpg, map[string]*_bintree_t{
			}},
			"banshee.jpg": &_bintree_t{assets_images_banshee_jpg, map[string]*_bintree_t{
			}},
			"chuck.jpg": &_bintree_t{assets_images_chuck_jpg, map[string]*_bintree_t{
			}},
			"dexter.jpg": &_bintree_t{assets_images_dexter_jpg, map[string]*_bintree_t{
			}},
			"elementary.jpg": &_bintree_t{assets_images_elementary_jpg, map[string]*_bintree_t{
			}},
			"suits.jpg": &_bintree_t{assets_images_suits_jpg, map[string]*_bintree_t{
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			".DS_Store": &_bintree_t{assets_js_ds_store, map[string]*_bintree_t{
			}},
			"bootstrap.js": &_bintree_t{assets_js_bootstrap_js, map[string]*_bintree_t{
			}},
			"bootstrap.min.js": &_bintree_t{assets_js_bootstrap_min_js, map[string]*_bintree_t{
			}},
			"de.js": &_bintree_t{assets_js_de_js, map[string]*_bintree_t{
			}},
			"external": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{assets_js_external_ds_store, map[string]*_bintree_t{
				}},
				"bootstrap3-typeahead.min.js": &_bintree_t{assets_js_external_bootstrap3_typeahead_min_js, map[string]*_bintree_t{
				}},
				"circle-progress.js": &_bintree_t{assets_js_external_circle_progress_js, map[string]*_bintree_t{
				}},
				"dataTables.bootstrap.min.js": &_bintree_t{assets_js_external_datatables_bootstrap_min_js, map[string]*_bintree_t{
				}},
				"fullcalendar.min.js": &_bintree_t{assets_js_external_fullcalendar_min_js, map[string]*_bintree_t{
				}},
				"jquery.dataTables.min.js": &_bintree_t{assets_js_external_jquery_datatables_min_js, map[string]*_bintree_t{
				}},
				"jquery.dynatable.js": &_bintree_t{assets_js_external_jquery_dynatable_js, map[string]*_bintree_t{
				}},
				"moment.min.js": &_bintree_t{assets_js_external_moment_min_js, map[string]*_bintree_t{
				}},
				"ohsnap.js": &_bintree_t{assets_js_external_ohsnap_js, map[string]*_bintree_t{
				}},
			}},
			"jquery.knob.min.js": &_bintree_t{assets_js_jquery_knob_min_js, map[string]*_bintree_t{
			}},
			"main.js": &_bintree_t{assets_js_main_js, map[string]*_bintree_t{
			}},
			"vendor": &_bintree_t{nil, map[string]*_bintree_t{
				".DS_Store": &_bintree_t{assets_js_vendor_ds_store, map[string]*_bintree_t{
				}},
				"jquery.ba-throttle-debounce.min.js": &_bintree_t{assets_js_vendor_jquery_ba_throttle_debounce_min_js, map[string]*_bintree_t{
				}},
				"jquery.min.js": &_bintree_t{assets_js_vendor_jquery_min_js, map[string]*_bintree_t{
				}},
				"underscore-min.js": &_bintree_t{assets_js_vendor_underscore_min_js, map[string]*_bintree_t{
				}},
			}},
		}},
		"templates": &_bintree_t{nil, map[string]*_bintree_t{
			".DS_Store": &_bintree_t{assets_templates_ds_store, map[string]*_bintree_t{
			}},
			"_header.tmpl": &_bintree_t{assets_templates_header_tmpl, map[string]*_bintree_t{
			}},
			"_script.tmpl": &_bintree_t{assets_templates_script_tmpl, map[string]*_bintree_t{
			}},
			"addSeries.tmpl": &_bintree_t{assets_templates_addseries_tmpl, map[string]*_bintree_t{
			}},
			"calendar.tmpl": &_bintree_t{assets_templates_calendar_tmpl, map[string]*_bintree_t{
			}},
			"history.tmpl": &_bintree_t{assets_templates_history_tmpl, map[string]*_bintree_t{
			}},
			"index.tmpl": &_bintree_t{assets_templates_index_tmpl, map[string]*_bintree_t{
			}},
			"presets.tmpl": &_bintree_t{assets_templates_presets_tmpl, map[string]*_bintree_t{
			}},
			"settings.tmpl": &_bintree_t{assets_templates_settings_tmpl, map[string]*_bintree_t{
			}},
			"show.tmpl": &_bintree_t{assets_templates_show_tmpl, map[string]*_bintree_t{
			}},
			"shows.tmpl": &_bintree_t{assets_templates_shows_tmpl, map[string]*_bintree_t{
			}},
		}},
		"translations": &_bintree_t{nil, map[string]*_bintree_t{
			".DS_Store": &_bintree_t{assets_translations_ds_store, map[string]*_bintree_t{
			}},
			"de-DE.all.json": &_bintree_t{assets_translations_de_de_all_json, map[string]*_bintree_t{
			}},
			"en-US.all.json": &_bintree_t{assets_translations_en_us_all_json, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

