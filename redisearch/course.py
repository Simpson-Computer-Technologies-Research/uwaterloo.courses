from redis_om import Field, HashModel

class Course(HashModel):
    components: str = Field(index=True)
    description: str = Field(index=True)
    id: str = Field(index=True)
    name: str = Field(index=True)
    pre_requisites: str = Field(index=True)
    title: str = Field(index=True)
    unit: str = Field(index=True)
