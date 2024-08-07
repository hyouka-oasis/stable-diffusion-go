interface OverrideSettings {
    sd_model_checkpoint?: string;
    sd_vae?: string;
}

interface Config {
    [key: string]: {
        override_settings?: OverrideSettings
        rol?: string;
        sampler_index?: string;
        width?: number;
        height?: number;
        steps?: number;
        seed?: number;
        cfg_scale?: number;
    };
}

const config: Config & {
    rol?: string
} = {
    "古色古香的人物形象": {
        override_settings: {
            sd_model_checkpoint: "GhostMix_V2.0"
        },
        width: 1024,
        height: 1024,
        sampler_index: "DPM++ 2S a Karras",
        steps: 35,
        seed: 2274007149,
        cfg_scale: 6,
        rol: "<lora:Manhuanan_20230827182922:0.7>,<lora:Imperial Water V2_2.0:0.7>,<lora:What a feeling as if in the mortal world {fantasy action lora}_1.0:0.7>,"
    }
};
