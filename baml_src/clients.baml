// Learn more about clients at https://docs.boundaryml.com/docs/snippets/clients/overview

client<llm> TogetherDeepSeekR1 {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "deepseek-ai/DeepSeek-R1"
    api_key env.TOGETHER_API_KEY
  }
}


client<llm> DeepSeekChat {
  provider "openai-generic"
  options {
    base_url "https://api.deepseek.com"
    model "deepseek-chat"
    api_key env.DEEPSEEK_API_KEY
  }
}
