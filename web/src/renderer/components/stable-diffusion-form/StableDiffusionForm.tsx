import { ProForm, ProFormDigit, ProFormSelect, ProFormTextArea } from "@ant-design/pro-components";
import { Form } from "antd";
import { stableDiffusionApi } from "renderer/api";

const StableDiffusionForm = () => {
    return (
        <>
            <ProForm.Group title={"模型配置"}>
                <ProFormSelect
                    width="sm"
                    name={[ "stableDiffusionConfig", "override_settings", "sd_model_checkpoint" ]}
                    label="模型/ckpt"
                    placeholder="请选择模型/ckpt"
                    rules={[ { required: true, message: '请选择模型/ckpt' } ]}
                    request={async () => {
                        const data = await stableDiffusionApi.getSdModels();
                        return data.list?.map(item => ({ label: item.title, value: item.model_name }));
                    }}
                />
                <ProFormSelect
                    width="sm"
                    name={[ "stableDiffusionConfig", "override_settings", "sd_vae" ]}
                    label="模型VAE"
                    placeholder="请选择模型VAE"
                    rules={[ { required: true, message: '请选择模型VAE' } ]}
                    request={async () => {
                        const data = await stableDiffusionApi.getSdVae();
                        return data.list?.map(item => ({ label: item.model_name, value: item.model_name }));
                    }}
                />
                <ProFormDigit
                    width="sm"
                    name={[ "stableDiffusionConfig", "clip_skip" ]}
                    label="Clip 跳过层"
                    placeholder="请填写Clip 跳过层"
                    min={1}
                    rules={[ { required: true, message: '请填写Clip 跳过层' } ]}
                />
            </ProForm.Group>
            <ProForm.Group title={"提示词配置"}>
                <ProFormTextArea
                    width="md"
                    name={[ "stableDiffusionConfig", "prompt" ]}
                    label="正向提示词/prompt"
                    placeholder="请填写正向提示词"
                />
                <ProFormTextArea
                    width="md"
                    name={[ "stableDiffusionConfig", "negative_prompt" ]}
                    label="反向提示词/negative_prompt"
                    placeholder="请填写反向提示词"
                />
            </ProForm.Group>
            <ProForm.Group title={"生成参数配置"}>
                <ProFormSelect
                    width="sm"
                    name={[ "stableDiffusionConfig", "sampler_name" ]}
                    label="采样器/Sampling method"
                    placeholder="请选择采样器/Sampling method"
                    rules={[ { required: true, message: '请选择采样器/Sampling method' } ]}
                    request={async () => {
                        const data = await stableDiffusionApi.getSamplers();
                        return data.list?.map(item => ({ label: item.name, value: item.name }));
                    }}
                />
                <ProFormSelect
                    width="sm"
                    name={[ "stableDiffusionConfig", "scheduler" ]}
                    label="调度类型/Schedule type"
                    placeholder="请选择调度类型/Schedule type"
                    rules={[ { required: true, message: '请选择调度类型/Schedule type' } ]}
                    request={async () => {
                        const data = await stableDiffusionApi.getSchedulers();
                        return data.list?.map(item => ({ label: item.label, value: item.name }));
                    }}
                />
                <ProFormDigit
                    width="sm"
                    name={[ "stableDiffusionConfig", "steps" ]}
                    label="采样步数/Sampling steps"
                    placeholder="请填写采样步数"
                    min={1}
                    rules={[ { required: true, message: '请填写采样步数' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "width" ]}
                    label="图片宽度"
                    placeholder="请填写图片宽度"
                    rules={[ { required: true, message: '请填写图片宽度' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "height" ]}
                    label="图片高度"
                    placeholder="请填写图片高度"
                    rules={[ { required: true, message: '请填写图片高度' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "n_iter" ]}
                    label="生成次数/Batch count"
                    placeholder="请填写生成次数"
                    min={1}
                    rules={[ { required: true, message: '请填写生成次数' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "batch_size" ]}
                    label="生成数量/batch_size"
                    min={1}
                    placeholder="请填写生成数量"
                    rules={[ { required: true, message: '请填写生成数量' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "cfg_scale" ]}
                    label="提示词引导系数/cfg_scale"
                    placeholder="请填写提示词引导系数"
                    min={1}
                    rules={[ { required: true, message: '请填写提示词引导系数' } ]}
                />
                <ProFormDigit
                    width="md"
                    name={[ "stableDiffusionConfig", "seed" ]}
                    label="随机数种子/seed"
                    placeholder="请填写随机数种子"
                    min={-99}
                    rules={[ { required: true, message: '请填写随机数种子' } ]}
                />

            </ProForm.Group>
            <ProForm.Group title={"高清修复配置"}>
                <ProFormSelect
                    width="sm"
                    name={[ "stableDiffusionConfig", "enable_hr" ]}
                    label="是否开启高清修复"
                    placeholder="请选择是否开启高清修复"
                    rules={[ { required: true, message: '请选择是否开启高清修复' } ]}
                    options={[
                        {
                            label: "是",
                            // @ts-ignore
                            value: true,
                        },
                        {
                            label: "否",
                            // @ts-ignore
                            value: false,
                        }
                    ]}
                />
                <Form.Item noStyle shouldUpdate>
                    {
                        ({ getFieldValue }) => {
                            const enableHr = getFieldValue([ "stableDiffusionConfig", "enable_hr" ]);
                            if (!enableHr) return null;
                            return (
                                <ProFormSelect
                                    width="sm"
                                    name={[ "stableDiffusionConfig", "hr_upscaler" ]}
                                    label="高清算法"
                                    placeholder="请选择高清算法"
                                    rules={[ { required: true, message: '请选择高清算法' } ]}
                                    request={async () => {
                                        const data = await stableDiffusionApi.getUpscalers();
                                        return data.list?.map(item => ({ label: item.name, value: item.name }));
                                    }}
                                />
                            );
                        }
                    }
                </Form.Item>
                <Form.Item noStyle shouldUpdate>
                    {
                        ({ getFieldValue }) => {
                            const enableHr = getFieldValue([ "stableDiffusionConfig", "enable_hr" ]);
                            if (!enableHr) return null;
                            return (
                                <ProFormDigit
                                    width="sm"
                                    name={[ "stableDiffusionConfig", "hr_second_pass_steps" ]}
                                    label="高分辨率采样步数/Hires steps"
                                    placeholder="请填写高分辨率采样步数/Hires steps"
                                    max={150}
                                    rules={[ { required: true, message: '请填写高分辨率采样步数/Hires steps' } ]}
                                />
                            );
                        }
                    }
                </Form.Item>
                <Form.Item noStyle shouldUpdate>
                    {
                        ({ getFieldValue }) => {
                            const enableHr = getFieldValue([ "stableDiffusionConfig", "enable_hr" ]);
                            if (!enableHr) return null;
                            return (
                                <ProFormDigit
                                    width="sm"
                                    name={[ "stableDiffusionConfig", "denoising_strength" ]}
                                    label="重绘强度/Denoising strength"
                                    placeholder="请填写重绘强度/Denoising strength"
                                    max={1}
                                    rules={[ { required: true, message: '请填写重绘强度/Denoising strength' } ]}
                                />
                            );
                        }
                    }
                </Form.Item>
                <Form.Item noStyle shouldUpdate>
                    {
                        ({ getFieldValue }) => {
                            const enableHr = getFieldValue([ "stableDiffusionConfig", "enable_hr" ]);
                            if (!enableHr) return null;
                            return (
                                <ProFormDigit
                                    width="sm"
                                    name={[ "stableDiffusionConfig", "hr_scale" ]}
                                    label="放大倍率/Upscale by"
                                    placeholder="请填写放大倍率/Upscale by"
                                    min={1}
                                    max={4}
                                    rules={[ { required: true, message: '请填写放大倍率/Upscale by' } ]}
                                />
                            );
                        }
                    }
                </Form.Item>
            </ProForm.Group>
        </>
    );
};

export default StableDiffusionForm;
