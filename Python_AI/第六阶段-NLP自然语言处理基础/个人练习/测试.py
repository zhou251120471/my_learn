from torch import nn


class Embedding(nn.Module):
    def __init__(self, vocab_size, embedding_dim):
        super(Embedding, self).__init__()
        # vocab_size: 词表大小
        # embedding_dim: 嵌入维度
        self.vocab_size = vocab_size
        self.embedding_dim = embedding_dim
        self.embedding = nn.Embedding(vocab_size, embedding_dim)

    def forward(self, x):
        return self.luc