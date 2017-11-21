package configs

import (
   "github.com/larspensjo/config"
)


func Load() (map[string]string){

    var topic = make(map[string]string)

    var sectionName = "default"
 
    cfg, _ := config.ReadDefault("config.ini")

    if cfg.HasSection(sectionName) {
        section, err := cfg.SectionOptions(sectionName)
        if err == nil {
            for _, v := range section {
                options, err := cfg.String(sectionName, v)
                if err == nil {
                    topic[v] = options
                }
            }
        }
    }

    return topic;
}

