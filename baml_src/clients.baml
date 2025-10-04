// Learn more about clients at https://docs.boundaryml.com/docs/snippets/clients/overview

client<llm> TogetherDeepSeekR1 {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "deepseek-ai/DeepSeek-R1"
    api_key env.TOGETHER_API_KEY
  }
}

client<llm> TogetherDeepSeekV3_1 {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "deepseek-ai/DeepSeek-V3.1"
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

client<llm> TogetherOpenAiGptOss20b {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "openai/gpt-oss-20b"
    api_key env.TOGETHER_API_KEY
  }
}

client<llm> TogetherLlama3_3_70bTurbo {
  provider "openai-generic"
  options {
    base_url "https://api.together.xyz/v1"
    model "meta-llama/Llama-3.3-70B-Instruct-Turbo-Free"
    api_key env.TOGETHER_API_KEY
  }
}